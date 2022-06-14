
## API Reference
## Table of Content
* [User](#Users)
* [Subject](#Subjects)
* [Post](#Posts)
* [Comments](#Comments)
* [Counts](#Count)
### Users

#### Get all users
Need the SID 
```http
  GET /api/users
```

**Code : 200(OK)**
Return an array of all users

#### Get user

```http
  GET /api/user/${id}
```


#### Create user

```http
  Post /api/user
```

| Parameter       | Type     | Description                                   |
| :--------       | :------- | :--------------------------------             |
| `UUID`      | `string` | **Required**. name of the subject                |
| `profilepicture`      | `[]byte` | **Required**. Description of the Subject           |
| `username`         | `string` | **Required**. UUID of the Subject          |
| `Password`     | `string` | **Not Required**. Image of the subject          |
| `Email`      | `int` | **Required**. Is the subject nsfw or not           |
| `firstname`     | `string` | **Not Required**. Tags of the subject           |
| `Lastname`      | `string` | **Required**. name of the subject                |
| `RiotId`      | `string` | **Required**. Description of the Subject           |
| `birthdate`         | `string` | **Required**. UUID of the Subject          |
| `oauthtoken`     | `string` | **Not Required**. Image of the subject          |
| `Genre`      | `int` | **Required**. Is the subject nsfw or not           |
| `Role`     | `string` | **Not Required**. Tags of the subject           |
| `Title`      | `string` | **Required**. name of the subject                |
| `Bio`      | `string` | **Required**. Description of the Subject           |
| `premium`         | `int` | **Required**. UUID of the Subject          |
| `follows`     | `string` | **Not Required**. Image of the subject          |

#### *Search User*
```http
  GET /api/user/search/:word
``` 

**Code : 200 (OK)**

#### *Get user by username*
```http
  GET /api/user/by-username/:username
```
**Code : 200 (OK)**

*Content example if we enter Shouyou instead of username*
```
{
    "uuid": "408b95ce-adfe-45e8-b657-b66ed81d76ec",
    "profilepicture": base64
    "username": "Shouyou",
    "password": "",
    "email": "eoazji@mgi.com",
    "firstname": "zeza",
    "lastname": "ezez",
    "riotid": "AAxd-MxCc_QfdYGj36pNkhMVYFHEsNYF-jaiNhq-EPa3C2c",
    "birthdate": "2022-06-01",
    "oauthtoken": "",
    "genre": "Other",
    "role": "",
    "title": "",
    "bio": "Un joueur de LoL \u0026 noir",
    "premium": 0,
    "follows": ""
}
```

#### Delete a user
```http
DELETE /apiuser/&{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to fetch |

**Code 200 (OK)**
*Content example*
```http
{
  "msg":"success"
}
```

### Subjects

#### Upvote
Need the SID

```http
GET /api/subject/:id/upvote
```

**Code : 200 (OK)**

*Content example*
``` 
{
  "msg":"success, 
}
```
#### Count 
Don't need the SID

```http
GET /api/subject/:id/count
```
**Code : 200 (OK)**

*Content example for a subject with only one UpVote and no DownVote*
``` 
{
    "UpVote": 1,
    "DownVote": 0
} 

```

#### Downvote
Need the SID

```http
GET /api/subject/:id/downvote
```
**Code : 200 (OK)**

*Content example DownVote*
```
{
    "msg": "success",
}
```
#### Vote
Need the SID

```http
GET /api/subject/:id/vote
```
**Code : 200 (OK)**

*Content example for a subject with only one UpVote and no DownVote*
```
{
    "downvote": false,
    "upvote": true
}
```
#### Search subject
Search a subject who contains the word 
```http
GET /api/subject/search/:word
```
**Code : 200 (OK)**

*Content example if we try with the word pirate*
```

  [{
    "id": "19",
    "title": "PIRATE",
    "description": "JOHNNY",
    "nsfw": 1,
    "image": Return the image in base64
    "tags": "#JVAIS TE manger",
    "upvotes": "#062568d4-fc0d-4e61-9685-4f3a895306a7",
    "downvotes": "",
    "publishdate": "10/6/2022",
    "lastpostdate": "2022-06-11 12:38:43",
    "allposts": "#17#19",
    "owner": "408b95ce-adfe-45e8-b657-b66ed81d76ec"
}, {
    "id": "22",
    "title": "La famille pirates",
    "description": "un tres bon dessin animé",
    "nsfw": 0,
    "image":Return the image in base64
    "tags": "#Pirate #Famille #DessinAnime",
    "upvotes": "",
    "downvotes": "",
    "publishdate": "11/6/2022",
    "lastpostdate": "",
    "allposts": "",
    "owner": "408b95ce-adfe-45e8-b657-b66ed81d76ec"
}]

```
#### Get a random subject
Don't requiert SID

Get a random subject 
```http
GET /api/subject/GetNBSubject/:nb 
```

**Code : 200 (OK)**

*Content example if we try with 2 as nb*
```
[{
    "id": "22",
    "title": "La famille pirates",
    "description": "un tres bon dessin animé",
    "nsfw": 0,
    "image": Return the image in base64
    "tags": "#Pirate #Famille #DessinAnime",
    "upvotes": "",
    "downvotes": "",
    "publishdate": "11/6/2022",
    "lastpostdate": "",
    "allposts": "",
    "owner": "408b95ce-adfe-45e8-b657-b66ed81d76ec"
}, {
    "id": "19",
    "title": "PIRATE",
    "description": "JOHNNY",
    "nsfw": 1,
    "image": Return the image in base64
    "tags": "#JVAIS TE manger",
    "upvotes": "#062568d4-fc0d-4e61-9685-4f3a895306a7",
    "downvotes": "",
    "publishdate": "10/6/2022",
    "lastpostdate": "2022-06-11 12:38:43",
    "allposts": "#17#19",
    "owner": "408b95ce-adfe-45e8-b657-b66ed81d76ec"
}]
```

#### Get the last subject update 
Don't requiert SID 
```http
GET /api/subject/GetLastSubjectUpdate/:nb
```
```
[{
    "id": "24",
    "title": "Trackmania",
    "description": "Vroum Vroum la voiture",
    "nsfw": 1,
    "image": Return the image in base64
    "tags": "#Pilote #Voiture #F1 #Vitesse",
    "upvote

}, {{
    "id": "23",
    "title": "Foot 2 Rue",
    "description": "Quand t aime le foot et que tu viens de la rue",
    "nsfw": 0,
    "image": return the image in base64
    "tags": "#Style #Foot #Rue #93",
    "upvotes": "",
    "downvotes": "",
    "publishdate": "11/6/2022",
    "lastpostdate": "2022-06-11 12:39:24",
    "allposts": "#20",
    "owner": "408b95ce-adfe-45e8-b657-b66ed81d76ec"
    }]
```
#### Get all Subjects

Show all subject.
```http
  GET /api/subjects
```

**Code : 200 (OK)**
Return all Subject 

#### Create Subject 
Create a subject.
```http
  Post /api/subject
```
| Parameter       | Type     | Description                                   |
| :--------       | :------- | :--------------------------------             |
| `title`      | `string` | **Required**. name of the subject                |
| `Description`      | `string` | **Required**. Description of the Subject           |
| `UUID`         | `string` | **Required**. UUID of the Subject          |
| `imagedata`     | `string` | **Not Required**. Image of the subject          |
| `nsfw`      | `int` | **Required**. Is the subject nsfw or not           |
| `tags`     | `string` | **Not Required**. Tags of the subject           |


#### Get subject
Show a subject.
```http
  GET /api/subject/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of subject to fetch |


#### Delete Subject
Delete a subject.
```http
  Delete /api/subject/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of subject to fetch |


**Code : 200 (OK)**

*Content example*
```http
{
    "msg": "success",
}
``` 

#### PUT Subject
Modify a subject
```http
  put /api/subject/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of subject to fetch |

**Code : 200 (OK)**

*Content example*
```http
{
    "msg": "success",
}
``` 



### Posts

#### Create post 
Create a post.
```http
  Post /api/posts
```

| Parameter       | Type     | Description                                   |
| :--------       | :------- | :--------------------------------             |
| `title`      | `string` | **Required**. name of the post                |
| `Description`      | `string` | **Required**. Description of the post           |
| `nsfw`      | `int` | **Required**. Is the post nsfw or not           |
| `tags`     | `string` | **Not Required**. Tags of the post           |
| `parent`     | `string` | **Not Required**. Parents of the post           |
| `UUID`     | `string` | **Not Required**. UUID of the post           |
| `image`     | `string` | **Not Required**. image of the subject           |

#### Get all posts

```http
  GET /api/posts
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get post
Show a post.
```http
  GET /api/post/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of subject to fetch |


#### Delete post
Delete a post.
```http
  Delete /api/post/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of subject to fetch |


#### PUT post
Modify a post
```http
  put /api/post/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of subject to fetch |

#### upvote

```http
GET /api/post/:id/upvote
```
**Code : 200 (OK)**

*Content example upvote*
```
{
    "msg": "success",
}
```

#### DownVote

```http
GET /api/post/:id/downvote
```
**Code : 200 (OK)**

*Content example DownVote*
```
{
    "msg": "success",
}
```

#### Count

```http
GET /api/post/:id/count
```
**Code : 200 (OK)**

*Content example for a post with only one UpVote and no DownVote*
``` 
{
    "UpVote": 1,
    "DownVote": 0
} 

```

#### Vote
Need the SID

```http
GET /api/post/:id/vote
```
**Code : 200 (OK)**

*Content example for a post with only one UpVote and no DownVote*
```
{
    "downvote": false,
    "upvote": true
}
```
#### Search Post 
```http
GET /api/post/search/:word
```
**Code : 200(OK)**

#### Get a random Post
Don't requiert SID

Get a random Post 
```http
GET /api/Post/GetNBPost/:nb 
```
**Code : 200(OK)**

#### Get last post

Don't requiert SID

Get the last post

```http
GET /api/post/GetLastPost/:nb 
```

**Code : 200(OK)**

#### GetPostBySubject

Don't requiert SID 

Get a post by using a subject id

```http
GET /api/post/GetPostBySubject/:id
```
**Code : 200(OK)**


### Comments

Show all comment.
```http
  GET /api/comments
```

**Code : 200 (OK)**

Return all Comments

####  Get comment by id 


Show all subject.
```http
  GET /api/comment/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of subject to fetch |

#### Like a comment 
Need the SID
```http 
GET /api/comment/:id/upvote
```
**Code : 200(OK)**

*Content example*
```
{
  "msg":"success"
}
```
#### Dislike a comment 
Need the SID
```http 
GET /api/comment/:id/downvote
```
**Code : 200(OK)**

*Content example*
```
{
  "msg":"success"
}
```

#### count like/dislike on a comment 
Don't need the SID
```http 
GET /api/comment/:id/count 
```
**Code : 200(OK)**

*Content example for a comment with only one UpVote and no DownVote*
``` 
{
    "UpVote": 1,
    "DownVote": 0
} 

```

#### Vote on a Comment 
Need the SID

```http
GET /api/comment/:id/vote
```
**Code : 200 (OK)**

*Content example for a comment with only one UpVote and no DownVote*
```
{
    "downvote": false,
    "upvote": true
}
```

#### Get comment by post id

```http
GET /api/comment/:id/vote
```

**Code : 200(OK)**

*Content example for a post with 2 comment ( here the post with the id 17)*
```http
[{
    "id": "13",
    "owner": "062568d4-fc0d-4e61-9685-4f3a895306a7",
    "content": "sqd",
    "upvotes": "",
    "downvotes": "",
    "parent": "17",
    "publishDate": "2022-06-13 16:54:41.455100"
}, {
    "id": "14",
    "owner": "062568d4-fc0d-4e61-9685-4f3a895306a7",
    "content": "qsd",
    "upvotes": "",
    "downvotes": "",
    "parent": "17",
    "publishDate": "2022-06-13 16:54:47.618512"
}]
```
### Count

#### Count the number of user, subject, userconnected ( session ) and number of post 
```http
GET /api/count
```

**Code : 200(OK)**
```http
{
    "Session": 2,
    "Subject": 7,
    "Post": 20,
    "User": 16
}
```

#### Count the number of User

```http
GET /api/count/user
```

**Code : 200(OK)**
```http
{
    "Nb": 16
}
```

#### Count the number of Post

```http
GET /api/count/post
```

**Code : 200(OK)**
```http
{
    "Nb": 20,
}
```

#### Count the number of subject
```http
GET /api/count/subject
```

**Code : 200(OK)**
```http
{
    "Nb": 7,

}
```
#### Count the number of Session ( User connected)
```http
GET /api/count/session
```

**Code : 200(OK)**
```http
{
    "Nb": 2,

}
```




