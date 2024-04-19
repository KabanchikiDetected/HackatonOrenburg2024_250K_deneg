from django.core.exceptions import BadRequest
from rest_framework.serializers import ListSerializer

from . import models
from . import serializers


class PostService:
    @staticmethod
    def get_all_by_user_id(user_id: str) -> ListSerializer[serializers.PostSerializer]:
        posts = models.PostModel.objects.filter(
            author_id=user_id
        )

        serializer = serializers.PostSerializer(
            posts, many=True
        )
        
        return serializer

    @staticmethod
    def create(post_data: dict) -> serializers.PostSerializer:
        serializer = serializers.PostSerializer(data=post_data)
        
        if not serializer.is_valid():
            raise BadRequest(serializer.error_messages)
        
        serializer.save()
        
        return serializer
    
    @staticmethod
    def update(post_id, user_id, new_post_data: dict) -> serializers.PostSerializer:
        post = PostService._get_by_id_and_user_id(post_id, user_id)
        
        serializer = serializers.PostSerializer(post, data=new_post_data)
        
        if not serializer.is_valid():
            raise BadRequest(serializer.errors)
        
        serializer.save()
        
        return serializer
    
    @staticmethod
    def delete(post_id: int, user_id: str):
        post = PostService._get_by_id_and_user_id(post_id, user_id)
        
        post.delete()
        
        serializer = serializers.EmptySerializer()
        serializer.data = "Post deleted"
        
        return serializer
    
    @staticmethod
    def get_by_id(post_id: int):
        serializer = serializers.PostSerializer(
            PostService._get_by_id(post_id)
        )
        
        return serializer
        
    @staticmethod
    def _get_by_id(post_id: int):
        try:
            post = models.PostModel.objects.get(
                pk=post_id
            )
        except models.PostModel.DoesNotExist:
            raise BadRequest(f"No post with id {post_id}")
        
        return post
    
    @staticmethod
    def get_by_id_and_user_id(post_id: int, user_id: str):
        serializer = serializers.PostSerializer(
            PostService._get_by_id_and_user_id(post_id, user_id)
        )

        return serializer
    
    @staticmethod
    def _get_by_id_and_user_id(post_id: int, user_id: str):
        try:
            post = models.PostModel.objects.get(
                pk=post_id,
                author_id=user_id
            )
        except models.PostModel.DoesNotExist:
            raise BadRequest(f"No post with id {post_id} by user with id {user_id}")
        
        return post


class LikePostService:
    @staticmethod
    def like(post_id: int, user_id: str):
        post = PostService._get_by_id(post_id)
        
        serializer = serializers.EmptySerializer()

        try:
            like = models.LikesToPostModel.objects.get(
                post_id=post_id,
                user_id=user_id
            )
            like.delete()
            
            post.likes -= 1
            post.save()
            
            serializer.data = "You removed the like"
        except models.LikesToPostModel.DoesNotExist:
            models.LikesToPostModel.objects.create(
                post_id=post_id,
                user_id=user_id
            )
            
            post.likes += 1
            post.save()
            
            serializer.data = "You have put a like"
            
        return serializer
    
    @staticmethod
    def is_user_like(post_id: int, user_id: str):
        post = PostService._get_by_id(post_id)

        serializer = serializers.EmptySerializer()

        try:
            models.LikesToPostModel.objects.get(
                post_id=post_id,
                user_id=user_id
            )
            
            serializer.data = True
        except models.LikesToPostModel.DoesNotExist:
            serializer.data = False
        
        return serializer


class PostImageService:
    def add_image(post_id: int, user_id: str, image_data: dict):
        post = PostService._get_by_id_and_user_id(post_id, user_id)
        
        serializer = serializers.PostImageSerializer(data=image_data)

        if not serializer.is_valid():
            print(serializer.error_messages)
            raise BadRequest(serializer.errors)
        
        serializer.save()
        
        serializer = serializers.EmptySerializer()
        serializer.data = "Image added"
        
        return serializer
