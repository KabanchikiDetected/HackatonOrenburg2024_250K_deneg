import json
from drf_spectacular.utils import extend_schema_field
from rest_framework import serializers

from . import models


class EmptySerializer:
    def __init__(self) -> None:
        self.data = None


class PostSerializer(serializers.ModelSerializer):
    images = serializers.SerializerMethodField()
    hashtags = serializers.SerializerMethodField()

    class Meta:
        model = models.PostModel
        fields = (
            'id', 'title', 'content', 'author_id', 'created_at', 'likes', 'images', 'hashtags'
        )
        read_only_fields = ('created_at', 'likes')
    
    @extend_schema_field(str)
    def get_images(self, obj):
        result = []
        
        for image in obj.images:
            result.append(image.image.url)
        
        return json.dumps(result)
    
    @extend_schema_field(str)
    def get_hashtags(self, obj):
        result = []
        
        for hashtag in obj.hashtags.all():
            result.append(hashtag.name)
        
        return json.dumps(result)


class PostImageSerializer(serializers.ModelSerializer):
    class Meta:
        model = models.PostImageModel
        fields = (
            "post_id", "image"
        )


class CommentSerializer(serializers.ModelSerializer):
    class Meta:
        model = models.CommentModel
        fields = (
            "id", "post_id", "user_id", "comment", "created_at"
        )
