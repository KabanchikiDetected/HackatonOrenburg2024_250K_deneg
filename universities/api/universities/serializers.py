import json
from drf_spectacular.utils import extend_schema_field
from rest_framework import serializers

from . import models


class EmptySerializer:
    def __init__(self) -> None:
        self.data = None
