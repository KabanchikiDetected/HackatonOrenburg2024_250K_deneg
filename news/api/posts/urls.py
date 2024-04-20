from django.contrib import admin
from django.urls import path, include

from . import controllers


urlpatterns = [
    path('', controllers.PostController.as_view()),
    path('<int:post_id>/like/', controllers.LikePostController.as_view()),
    path('<int:post_id>/image/', controllers.PostImageController.as_view()),
    path('<int:post_id>/comments/', controllers.CommentsPostController.as_view()),
    
    path("feed/", controllers.FeedPostController.as_view()),
]
