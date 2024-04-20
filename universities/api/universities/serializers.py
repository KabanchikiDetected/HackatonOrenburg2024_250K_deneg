import json
from drf_spectacular.utils import extend_schema_field
from rest_framework import serializers

from . import models


class EmptySerializer:
    def __init__(self) -> None:
        self.data = None


class UniversitySerializer(serializers.ModelSerializer):
    images = serializers.SerializerMethodField()

    class Meta:
        fields = "__all__"
        read_only_fields = ("images", )
        model = models.UniversityModel
        
    def get_images(self, obj):
        result = []

        for image in obj.images:
            result.append(image.image.url)
        
        return json.dumps(result)
        

class DepartmentSerializer(serializers.ModelSerializer):
    class Meta:
        fields = "__all__"
        model = models.DepartmentModel
        
        
class GroupSerializer(serializers.ModelSerializer):
    class Meta:
        fields = "__all__"
        model = models.GroupModel


class UserToGroupSerializer(serializers.ModelSerializer):
    class Meta:
        model = models.UserToGroupModel
        fields = "__all__"


class UniversityImageSerialzier(serializers.ModelSerializer):
    class Meta:
        model = models.UniversityImageModel
        fields = (
            "university_id", "image"
        )
