

1. Create namespace `oc new-project ex-dra-driver`
2. Install Scenario - `kustomize build examples/<scenario> | oc apply -f -`
3. Uninstall Scenario - `kustomize build examples/<scenario> | oc delete -f -`

The scenarios are:

1. `two-pods-two-distinct-nx-gzip` - Two pods, one container each, Each container asking for access to the nx-gzip
- Connect to one of the running pods, and you should see:
```
❯ oc get pods -n ex-dra-driver                      
NAME                                 READY   STATUS    RESTARTS   AGE
ex-dra-driver-555db57b69-gswt8       1/1     Running   0          15s
ex-dra-driver-sec-7967c5f977-xj2cm   0/1     Pending   0          15s
❯ oc rsh -nex-dra-driver  pod/ex-dra-driver-555db57b69-gswt8
sh-5.1$ ls -al /dev/crypto/nx-gzip 
crw-rw-rw-. 1 1000740023 root 243, 0 Aug 12 12:03 /dev/crypto/nx-gzip
```

2. `structured-parameters-one-pod` demonstrates the simple use-case. You should see `/dev/crypto/nx-gzip` mounted in the Pod.

3. Run the example demo test using `oc apply -f example/demo/nx-test.yaml`

These examples are inspired by https://github.com/kubernetes-sigs/dra-example-driver/tree/main/demo