## Docker build

    docker build . --file Dockerfile --tag marianferenc/project_atlas_cluster_processor_service:latest
	
## Kubernetes deployment

	helm upgrade --install atlas-api-service charts/cluster_processor_service --namespace project-atlas-system
	
## Kubernetes undeployment

    helm del atlas-api-service --namespace project-atlas-system
    
## Forward deployed application to localhost

    kubectl port-forward svc/atlas-api-service 3000:3000 --namespace project-atlas-system
    
