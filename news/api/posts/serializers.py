import json
from rest_framework import serializers

from . import models


class EmptySerializer:
    def __init__(self) -> None:
        self.data = None


class PostSerializer(serializers.ModelSerializer):
    images = serializers.SerializerMethodField()

    class Meta:
        model = models.PostModel
        fields = (
            'id', 'title', 'content', 'author_id', 'created_at', 'likes', 'images'
        )
        
    def get_images(self, obj):
        images = []
        
        for image in obj.images:
            images.append(
                image.image.url
            )
        
        return json.dumps(images)


class PostImageSerializer(serializers.ModelSerializer):
    class Meta:
        model = models.PostImageModel
        fields = (
            "post_id", "image"
        )
