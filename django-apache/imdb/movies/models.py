from django.db import models

# Create your models here.
class Movie(models.Model):
    name = models.CharField(max_length = 200)
    release_date = models.DateTimeField()
    description = models.CharField(max_length =500)

    @classmethod
    def create(cls, name, release_date, description):
        return cls(name=name, release_date=release_date, description=description)