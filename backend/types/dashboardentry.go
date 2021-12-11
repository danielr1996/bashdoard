package types

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type DashboardEntry struct {
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	DisplayName     string `json:"displayname"`
	Url             string `json:"url"`
	FaviconLocation string `json:"faviconLocation"`
}

/*func (r DashboardEntry) equals(d DashboardEntry) bool {
	return r.Name == d.Name &&
		r.Namespace == d.Namespace &&
		r.DisplayName == d.DisplayName &&
		r.Url == d.DisplayName &&
		r.FaviconLocation == d.FaviconLocation
}*/

func ToDashboardEntry(obj interface{}) DashboardEntry {
	u := obj.(*unstructured.Unstructured)
	name := u.GetName()
	namespace := u.GetNamespace()
	url, found, err := unstructured.NestedString(u.Object, "spec", "url")
	if !found || err != nil {
		panic(err.Error())
	}
	faviconLocation, found, err := unstructured.NestedString(u.Object, "spec", "faviconLocation")
	if !found || err != nil {
		panic(err.Error())
	}
	displayname, found, err := unstructured.NestedString(u.Object, "spec", "name")
	if !found || err != nil {
		panic(err.Error())
	}
	return DashboardEntry{Name: name, Namespace: namespace, DisplayName: displayname, Url: url, FaviconLocation: faviconLocation}
}
