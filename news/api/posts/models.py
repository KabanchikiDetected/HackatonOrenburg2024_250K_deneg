from typing import Self
from django.db import models
import random
import time


class UnixTimestampManager(models.Manager):
    def create(self, **kwargs):
        kwargs['created_at'] = str(time.time())
        return super().create(**kwargs)


def get_random_number():
    return random.randint(10000, 99999)


class Hashtag(models.Model):
    name = models.CharField(max_length=100, unique=True)

    def __str__(self):
        return '#' + self.name


class PostModel(models.Model):
    title = models.CharField(
        "Название", max_length=512
    )    
    content = models.TextField(
        "Содержимое"
    )
    author_id = models.CharField(
        "Id автора", max_length=128
    )

    likes = models.IntegerField(
        "Кол-во лайков", default=0
    )
    
    hashtags = models.ManyToManyField(Hashtag, related_name='posts')

    created_at = models.CharField(
        "Дата создания", max_length=128,
        blank=True, null=True
    )
    
    objects = UnixTimestampManager()

    class Meta:
        db_table = "post"
        verbose_name = "Пост"

    @property
    def images(self):
        images = PostImageModel.objects.filter(
            post_id=self.pk
        )
        
        return images
    
    def set_hashtags(self, hashtags: list[str]):
        self.hashtags.clear()

        for hashtag in hashtags:
            hashtag_object, status = Hashtag.objects.get_or_create(
                name=hashtag
            )
            
            self.hashtags.add(hashtag_object)
            
            
class CommentModel(models.Model):
    post_id = models.IntegerField(
        "Id пользователя"
    )
    user_id = models.CharField(
        "Id пользователя", max_length=128
    )
    comment = models.CharField(
        "Комментарий", max_length=1024
    )
    created_at = models.CharField(
        "Дата создания", max_length=128,
        blank=True, null=True
    )
    
    objects = UnixTimestampManager()
    
    class Meta:
        db_table = "comment"
        verbose_name = "Комментарий"


class LikesToPostModel(models.Model):
    user_id = models.CharField(
        "Id пользователя", max_length=128
    )
    post_id = models.IntegerField(
        "Id поста"
    )
    
    class Meta:
        db_table = "like_to_post"
        verbose_name = "Лайки к постам"


def post_image_upload_path(instance, filename):
    return f'post_image/{instance.post_id}/{get_random_number()}_{filename}'


class PostImageModel(models.Model):
    post_id = models.IntegerField(
        "Id поста"
    )
    image = models.ImageField(
        "Изображение",
        upload_to=post_image_upload_path,
        max_length=512,
    )
    
    class Meta:
        db_table = "post_image"
        verbose_name = "Лайки к постам"
