ctx = {
    resource_list.items = [
        {
            kind = "Deployment"
            apiVersion = "apps/v1"
            metadata.annotations = {
                "config.kubernetes.io/local-config" = "true"
            }
        }
    ]
    resource_list.functionConfig = {
        data.replicas = 5
        data.source = """\
params = ctx.resource_list.functionConfig.params
toMatch = params.toMatch
toAdd = params.toAdd
items = ctx?.resource_list?.items or []
[item | {
   if all key, value in toMatch {
      item.metadata.annotations[key] == value
   }:
      metadata.annotations: {**params.toAdd}
} for item in items]
"""
        params = {
            toMatch: {
                "config.kubernetes.io/local-config" = "true"
            }
            toAdd: {
                "configmanagement.gke.io/managed" = "disabled"
            }
        }
    }
}

items = (lambda {
  params = ctx.resource_list.functionConfig.params
  toMatch = params.toMatch
  toAdd = params.toAdd
  items = ctx?.resource_list?.items or []
  [item | {
     # If all annotations are matched, patch more annotations
     if all key, value in toMatch {
        item.metadata.annotations[key] == value
     }:
        metadata.annotations: {**params.toAdd}
  } for item in items]
})()
