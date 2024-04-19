from django.contrib import admin
from django.urls import path, include

from . import controllers


urlpatterns = [
    path("department/", controllers.DepartmentController.as_view()),
    path("university/", controllers.UniversityController.as_view()),
    path("group/", controllers.GroupController.as_view()),
]
