from django.contrib import admin
from django.urls import path, include

from . import controllers


urlpatterns = [
    path("university/", controllers.UniversitiesController.as_view()),
    path("university/<int:university_id>/", controllers.UniversityController.as_view()),
    path("university/<int:university_id>/add-image/", controllers.UniversityAddImageController.as_view()),
    
    path("university/<int:university_id>/department/", controllers.DepartmentsController.as_view()),
    path("university/<int:university_id>/department/<int:department_id>/", controllers.DepartmentController.as_view()),

    path("university/<int:university_id>/department/<int:department_id>/group/", controllers.GroupsController.as_view()),
    path("university/<int:university_id>/department/<int:department_id>/group/<int:group_id>/", controllers.GroupController.as_view()),
    
    path("university/<int:university_id>/department/<int:department_id>/group/<int:group_id>/requests/", controllers.GroupRequestsController.as_view()),
    path("university/<int:university_id>/department/<int:department_id>/group/<int:group_id>/requests/<int:request_id>/accept/", controllers.GroupRequestAcceptController.as_view()),
    path("university/<int:university_id>/department/<int:department_id>/group/<int:group_id>/requests/<int:request_id>/deny/", controllers.GroupRequestDenyController.as_view()),
    
    path("search/university/", controllers.UniversitySearchController.as_view()),
    path("search/department/", controllers.SearchDepartmentController.as_view()),
    path("search/group/", controllers.SearchGroupController.as_view()),
    
    path("my/university/", controllers.DeputyUniversityController.as_view()),
    path("my/group/", controllers.MyGroupController.as_view()),
    path("my/request/", controllers.MyRequestController.as_view()),
]
