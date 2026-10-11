# lab-hpa

Um hello world que escala sozinho com **HorizontalPodAutoscaler**. Feito para rodar no
[Developer Sandbox do OpenShift](https://developers.redhat.com/developer-sandbox), mas
os manifestos são Kubernetes puro (fora a `Route`).

```
app/    hello world em Go + Dockerfile
k8s/    app.yaml (Deployment, Service, Route), hpa.yaml, load.yaml (gerador de carga)
```

## Como o HPA decide

```
réplicas desejadas = ceil( réplicas atuais × CPU média atual / CPU alvo )
```

O alvo é **50% do request** de CPU. O request é 20m, então acima de ~10m de CPU
por pod ele cria pods (até 5); abaixo, remove (até 1), esperando 1 minuto de calma.

## Rodando

1. **Imagem**: troque `SEU_USUARIO` em `k8s/app.yaml` pelo seu usuário do Docker Hub.
   A imagem é publicada pelo workflow (`.github/workflows/ci.yml`) a cada push na `main`.
   Ou publique na mão:

   ```sh
   docker build -t docker.io/SEU_USUARIO/hello-hpa:latest app
   docker push docker.io/SEU_USUARIO/hello-hpa:latest
   ```

2. **Login no sandbox**: no console, canto superior direito → *Copy login command*.

   ```sh
   oc login --token=... --server=...
   ```

3. **Subir**:

   ```sh
   oc apply -f k8s/app.yaml -f k8s/hpa.yaml
   curl https://$(oc get route hello -o jsonpath='{.spec.host}')
   ```

4. **Ver escalar**: deixe um terminal assistindo e dispare a carga no outro.

   ```sh
   oc get hpa hello -w          # terminal 1
   oc create -f k8s/load.yaml   # terminal 2
   ```

   Em ~1 minuto a coluna `REPLICAS` sobe. A carga dura 3 minutos; depois disso,
   mais ~1 minuto e volta para 1.

5. **Limpar**:

   ```sh
   oc delete -f k8s/app.yaml -f k8s/hpa.yaml
   ```

## Se não escalar

- `TARGETS` mostra `<unknown>`: as métricas ainda não chegaram (espere 1 min) ou o
  pod está sem `resources.requests.cpu` — sem request, o HPA não tem base de cálculo.
- `oc describe hpa hello` mostra os eventos e o motivo de cada decisão.
