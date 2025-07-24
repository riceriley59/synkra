import json
import sys
from typing import Dict, Any, List


class Resource:
    def __init__(self, typ: str, name: str, props: Dict[str, Any]):
        self.type = typ
        self.name = name
        self.properties = props

    def to_dict(self) -> Dict[str, Any]:
        return {
            "type": self.type,
            "name": self.name,
            "properties": self.properties
        }


class Context:
    def __init__(self):
        self.resources: List[Resource] = []

    def register_resource(self, typ: str, name: str, props: Dict[str, Any]) -> Resource:
        resource = Resource(typ, name, props)
        self.resources.append(resource)
        self._send_to_engine(resource)
        return resource

    def export_resources(self, path: str):
        with open(path, "w") as f:
            json.dump([r.to_dict() for r in self.resources], f, indent=2)

    def _send_to_engine(self, resource: Resource):
        print("Registering resource:")
        print(json.dumps(resource.to_dict(), indent=2), file=sys.stdout)


# Global context — simulates the Pulumi runtime
global_context = Context()
