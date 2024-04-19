from django.core.exceptions import BadRequest
from rest_framework.request import Request
from rest_framework import permissions
from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework import status, generics
from drf_spectacular.utils import extend_schema, extend_schema_view, OpenApiParameter, OpenApiResponse, OpenApiExample

from . import docs
from . import services
from api import pagination
from . import serializers


@extend_schema_view(
    get=extend_schema(
        summary="Получить список постов пользователя",
        description="Метод возвращает все посты пользователя (пользователя получает по user_id из query параметров или из токена)",
        responses={
            status.HTTP_200_OK: serializers.PostSerializer(many=True)
        },
        parameters=[
            OpenApiParameter(
                name='user_id',
                location=OpenApiParameter.QUERY,
                description='Id пользователя для получения постов',
                required=False,
                type=str
            )
        ]
    ),
    post=extend_schema(
        summary="Создание нового поста",
        description="Создание нового поста от лица пользователя (пользователя получает по токену)",
        responses={
            status.HTTP_201_CREATED: serializers.PostSerializer
        },
        request=serializers.PostSerializer
    ),
    put=extend_schema(
        summary="Обновить пост",
        description="Обновляет пост по id, если он принадлежит пользователю (пользователя получает по токену)",
        responses={
            status.HTTP_200_OK: serializers.PostSerializer
        },
        request=docs.PutRequestPostSerialzier
    ),
    delete=extend_schema(
        summary="Удалить пост",
        description="Удаляет пост по id из query параметров, если он принадлежит пользователю (пользователя получает по токену)",
        parameters=[
            OpenApiParameter(
                name='post_id',
                location=OpenApiParameter.QUERY,
                description='Id поста для удаления',
                required=True,
                type=int
            ),
        ],
        responses={
            status.HTTP_200_OK: {
                "example": "Post deleted"    
            },
        }
    ),
)
class PostController(APIView):
    serializer_class = serializers.PostSerializer
    
    def get_permissions(self):
        if self.request.method == 'GET':
            return [permissions.AllowAny()]

        return super().get_permissions()    

    def get(self, request: Request):
        user_id = request.query_params.get("user_id") or request.user_data.get("id")
        
        if not user_id:
            raise BadRequest("No token or user_id in query params")
        
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


@extend_schema_view(
    get=extend_schema(
        summary="Стоит ли лайк на посте?",
        description="Метод проверяет стоит ли лайк у пользователя (берет по токену) на этом посте",
        responses={
            status.HTTP_200_OK: {
                "example": True
            }
        }
    ),
    post=extend_schema(
        summary="Поставить или урать лайк",
        description="Метод ставит лайк, если его небыло и убирает, если был",
        responses={
            status.HTTP_201_CREATED: {
                "example": ["You removed the like", "You have put a like"]
            }
        },
    )
)
class LikePostController(APIView):
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


@extend_schema_view(
    post=extend_schema(
        summary="Добавить фото к посту",
        description=(
            "Метод добавляет фото, если пост пинадлежит пользователю. "
            "(передавать post_id в body не обязательно, мне просто впадлу нормальную доку писать)"
        ),
        responses={
            status.HTTP_201_CREATED: serializers.PostImageSerializer
        },
        request=serializers.PostImageSerializer
    )
)
class PostImageController(APIView):
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

@extend_schema_view(
    get=extend_schema(
        summary="Лента новостей",
        description="Ендпоинт просто возвращает все посты от последнего к первому"
    )
)
class FeedPostController(generics.ListAPIView):
    queryset = services.PostService._get_all()
    serializer_class = serializers.PostSerializer
    pagination_class = pagination.StandartPagination
    permission_classes = (permissions.AllowAny, )
