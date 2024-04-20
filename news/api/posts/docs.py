from rest_framework import serializers


class PostRequestPostSerialzier(serializers.Serializer):
    title = serializers.CharField()
    content = serializers.CharField()
    hashtags = serializers.CharField()


class PutRequestPostSerialzier(serializers.Serializer):
    id = serializers.IntegerField()
    title = serializers.CharField()
    content = serializers.CharField()
    hashtags = serializers.CharField()

class PostRequestCommentPostSerialzier(serializers.Serializer):
    comment = serializers.CharField()
