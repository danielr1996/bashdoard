# bashdoard

![image](https://user-images.githubusercontent.com/6663726/174461547-1721f249-5df3-48d7-bb17-0d69ce21ce7c.png)

bashdoard is a cloud native dashboard application that provides links to all your self hosted apps in one place. 

## Usage
### Deployment
Deploy the dashboard app with docker compose

```
services:
  backend:
    build: ghcr.io/danielr1996/bashdoard-backend:latest
    ports:
      - "3040:3040"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
  frontend:
    build: ghcr.io/danielr1996/bashdoard-frontend:latest
    ports:
      - "8081:80"
    environment:
      API_URL: http://localhost:3040
```

### Providers

Currently the only supported provider is docker via container labels, but support for more providers like file or kubernetes
could be easily implemented

#### DockerProvider
```
services: 
  nginx:
    image: nginx
    labels:
      de.danielr1996.bashdoard.name: Nginx
      de.danielr1996.bashdoard.url: https://nginx.app.danielr1996.de
      de.danielr1996.bashdoard.icon: /favicon.ico
      de.danielr1996.bashdoard.id: nginx
  apache:
    image: apache
    labels:
      de.danielr1996.bashdoard.name: Apache
      de.danielr1996.bashdoard.url: https://apache.app.danielr1996.de
      de.danielr1996.bashdoard.icon: /favicon.png
      de.danielr1996.bashdoard.id: apache
```

## Roadmap
* Automatically read entries from Ingress, Ingressroute, Service, etc. 
* Add Tags
* Customizable Dashboard filtered by tags
* Rework UI
