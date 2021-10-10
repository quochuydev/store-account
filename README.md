1.  cd mynginx
    helm lint ./
    helm template store-account-nginx ./

2.  cd ..
    helm list
    helm install store-account-mynginx ./mynginx/
    kubectl describe deployment store-account-mynginx
    kubectl get all

3.  mkdir publish
    helm package ./mynginx -d ./publish/
    helm repo index .

- publish:
  helm repo add store-account-mynginx https://raw.githubusercontent.com/username/store-account-mynginx/master/
