from typing import Any
from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework import status

from . import services
from . import serializers
from api.pagination import StandartPagination


class BaseController(APIView):
    def perform_authentication(self, request):
        return None


class PostController(BaseController):
    serializer_class = serializers.PostSerializer

    def get(self, request: Request):
        user_id = request.user_data.get("id")
        
        serializer = services.PostService.get_all_by_user_id(user_id)
    
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )

    def post(self, request: Request):
        user_id = request.user_data.get("id")
        
        serializer = services.PostService.create({
            **request.data,
            "author_id": user_id
        })

        return Response(
            serializer.data,
            status.HTTP_201_CREATED
        )
    
    def put(self, request: Request):
        user_id = request.user_data.get("id")
        
        post_id = request.data.get("id")
        
        serializer = services.PostService.update(post_id, user_id, {
            "title": request.data.get("title", ""),
            "content": request.data.get("content", ""),
            "author_id": user_id
        })
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )
        
    def delete(self, request: Request):
        user_id = request.user_data.get("id")
        post_id = request.query_params.get("post_id")
        
        serializer = services.PostService.delete(post_id, user_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class LikePostController(BaseController):
    def get(self, request: Request, post_id: int):
        user_id = request.user_data.get("id")
        
        serializer = services.LikePostService.is_user_like(post_id, user_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )

    def post(self, request: Request, post_id: int):
        user_id = request.user_data.get("id")
        
        serializer = services.LikePostService.like(post_id, user_id)
        
        return Response(
            serializer.data,
            status.HTTP_201_CREATED
        )


class PostImageController(BaseController):
    def post(self, request: Request, post_id: int):
        user_id = request.user_data.get("id")
        
        serializer = services.PostImageService.add_image(post_id, user_id, {
            "image": request.FILES.get("image"),
            "post_id": post_id,
        })
        
        return Response(
            serializer.data,
            status.HTTP_201_CREATED
        )