ctx = {
    resource_list.items = [
        {
            kind = "Deployment"
            apiVersion = "apps/v1"
        }
    ]
    resource_list.functionConfig = {
        data.replicas = 5
        data.source = """\
setReplicas = lambda items, replicas {
   [item | {
      spec.replicas = replicas
   } if item.kind == "Deployment" and item.apiVersion == "apps/v1" else {} for item in items or {}]
}
setReplicas(ctx.resource_list?.items, ctx.resource_list?.functionConfig?.data?.replicas)
"""
    }
}

items = (lambda {
    setReplicas = lambda items, replicas {
       [item | {
          spec.replicas = replicas
       } if item.kind == "Deployment" and item.apiVersion == "apps/v1" else {} for item in items or {}]
    }
    setReplicas(ctx.resource_list?.items, ctx.resource_list?.functionConfig?.data?.replicas)
})()
