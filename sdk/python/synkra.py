import json
import sys
from typing import Dict, Any, List
import os
import grpc
import pingpong_pb2
import pingpong_pb2_grpc
import importlib.util


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


def register_resource(typ, name, props):
    global_context.register_resource(typ, name, props)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: synkra.py <user_file.py>")
        sys.exit(1)
    user_file = sys.argv[1]
    # Make register_resource available to the user file
    globals_dict = globals().copy()
    globals_dict["register_resource"] = register_resource
    # Import and execute the user file
    spec = importlib.util.spec_from_file_location("user_module", user_file)
    user_module = importlib.util.module_from_spec(spec)
    user_module.register_resource = register_resource
    sys.modules["user_module"] = user_module
    spec.loader.exec_module(user_module)
