from django.shortcuts import render
from rest_framework.views import APIView
from rest_framework.response import Response
from .models import Movie

# Create your views here.
class MovieController(APIView):

    def get(self, request):
        movies = Movie.objects.all().values('name', 'release_date', 'description')
        return Response(movies)
