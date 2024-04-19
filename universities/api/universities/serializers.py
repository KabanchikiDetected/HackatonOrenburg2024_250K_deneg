import json
from drf_spectacular.utils import extend_schema_field
from rest_framework import serializers

from . import models


class EmptySerializer:
    def __init__(self) -> None:
        self.data = None


class UniversitySerializer(serializers.ModelSerializer):
    class Meta:
        fields = "__all__"
        model = models.UniversityModel
        

class DepartmentSerializer(serializers.ModelSerializer):
    class Meta:
        fields = "__all__"
        model = models.DepartmentModel
        
        
class GroupSerializer(serializers.ModelSerializer):
    class Meta:
        fields = "__all__"
        model = models.GroupModel
