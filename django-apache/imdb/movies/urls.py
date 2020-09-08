from django.urls import path
from movies.views import MovieController

urlpatterns = [
    path('', MovieController.as_view())
]