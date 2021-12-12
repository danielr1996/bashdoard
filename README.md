# bashdoard

bashdoard is a kubernetes native dashboard application that provides links to all your self hosted apps in one place. 

## Usage
Install with helm
```
helm repo add danielr1996 https://danielr1996.github.io/k8s-charts
helm upgrade --install bashdoard danielr1996/bashdoard
```

Install `DashboardEntry` with kubectl
``` shell
cat <<EOF | kubectl apply -f - 
apiVersion: "bashdoard.danielr1996.de/v1alpha"
kind: DashboardEntry
metadata:
  name: traefik
spec:
  name: "Traefik"
  url: https://traefik.app.local.danielr1996.de
  faviconLocation: /dashboard/statics/icons/favicon.ico
EOF
```

## Roadmap
* Automatically read entries from Ingress, Ingressroute, Service, etc. 
* Add Tags
* Customizable Dashboard filtered by tags
* Rework UI
