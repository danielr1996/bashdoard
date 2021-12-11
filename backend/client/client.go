package client

import (
	"context"
	"danielr1996/bashboard/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"gopkg.in/antage/eventsource.v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

type DashboardEntryClient struct {
	informer cache.SharedIndexInformer
	es       eventsource.EventSource
	client   dynamic.Interface
}

var gvr = schema.GroupVersionResource{Group: "bashdoard.danielr1996.de", Version: "v1alpha", Resource: "dashboardentries"}

func New() *DashboardEntryClient {
	c := new(DashboardEntryClient)
	kubeCfg, err := rest.InClusterConfig()
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		kubeCfg, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	if err != nil {
		panic(err.Error())
	}

	client, err := dynamic.NewForConfig(kubeCfg)
	c.client = client
	if err != nil {
		panic(err.Error())
	}
	factory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(c.client, 0, corev1.NamespaceAll, nil)
	c.informer = factory.ForResource(gvr).Informer()

	c.es = eventsource.New(
		eventsource.DefaultSettings(),
		func(req *http.Request) [][]byte {
			return [][]byte{
				[]byte("Access-Control-Allow-Origin: *"),
			}
		},
	)
	return c
}

func (r *DashboardEntryClient) StartWatching(stopCh <-chan struct{}) {
	r.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			dbe := types.ToDashboardEntry(obj)
			fmt.Print("ADD: ")
			fmt.Println(dbe)
			b, err := json.Marshal(dbe)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
			r.es.SendEventMessage(string(b), "add", uuid.New().String())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			old := types.ToDashboardEntry(oldObj)
			new := types.ToDashboardEntry(newObj)
			fmt.Print("OLD: ")
			fmt.Println(old)
			fmt.Print("NEW: ")
			fmt.Println(new)
			b, err := json.Marshal(new)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
			r.es.SendEventMessage(string(b), "update", uuid.New().String())
		},
		DeleteFunc: func(obj interface{}) {
			dbe := types.ToDashboardEntry(obj)
			fmt.Print("DEL: ")
			fmt.Println(dbe)
			b, err := json.Marshal(dbe)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
			r.es.SendEventMessage(string(b), "delete", uuid.New().String())
		},
	})

	r.informer.Run(stopCh)
}

func (r *DashboardEntryClient) PushUpdates(stopCh <-chan struct{}) {
	http.Handle("/api/dashboardentries", r.es)
	http.HandleFunc("/api/sync", func(w http.ResponseWriter, req *http.Request) {
		list, err := r.client.Resource(gvr).Namespace("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		var entries []types.DashboardEntry
		for _, entry := range list.Items {
			entries = append(entries, types.ToDashboardEntry(&entry))
		}
		b, err := json.Marshal(entries)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, string(b))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	<-stopCh
	r.es.Close()
}
