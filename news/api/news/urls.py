from django.contrib import admin
from django.urls import path, include

from . import controllers


urlpatterns = [
    path('news', controllers.NewsController.as_view())
]
