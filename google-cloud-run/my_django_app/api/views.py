
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status

class ApiTest(APIView):
    def get(self, request):
        data={
            "greeting": "Say Hello!"
        }
        return Response(data, status=status.HTTP_200_OK, headers=None)