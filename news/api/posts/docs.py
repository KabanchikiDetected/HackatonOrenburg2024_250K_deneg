from rest_framework import serializers


class PutRequestPostSerialzier(serializers.Serializer):
    id = serializers.IntegerField()
    title = serializers.CharField()
    content = serializers.CharField()
