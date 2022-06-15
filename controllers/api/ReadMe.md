## API Reference

## Table of Contents
* [User](#users)
* [Subject](#subjects)
* [Post](#posts)
* [Comments](#comments)
* [Counts](#count)

### Users

#### Get all users
Need the SID 
```http
  GET /api/users
```

**Code : 200(OK)**
Return an array of all users

#### Get user
Don't need the SID
```http
  GET /api/user/${id}
```


#### Create user
Don't need the SID

```http
  Post /api/user
```

| Parameter       | Type     | Description                            |
| :--------       | :------- | :--------------------------------      |
| `profilepicture`| `[]byte` | **Not Required**. Image of the user    |
| `username`      | `string` | **Required**. Username of the user     |
| `Password`      | `string` | **Required**. Password of the user     |
| `Email`         | `string` | **Required**. Email of the user        |
| `firstname`     | `string` | **Required**. Firstname of the user    |
| `Lastname`      | `string` | **Required**. Lastname of the user     |
| `RiotId`        | `string` | **Not Required**. RiotId of the user   |
| `birthdate`     | `string` | **Required**. Birthdate of the user    |
| `Genre`         | `string` | **Required**. Genre of the user        |
| `Bio`           | `string` | **Not Required**. Bio of the user      |

#### *Search User*
Don't need the SID

**Code : 200(OK)**

```http
  GET /api/user/search/:word
``` 

**Code : 200 (OK)**

#### *Get user by username*
Don't need the SID
```http
  GET /api/user/by-username/:username
```
**Code : 200 (OK)**

*Content example if we enter Shouyou instead of username*
```json
{
    "uuid": "408b95ce-adfe-45e8-b657-b66ed81d76ec",
    "profilepicture": base64, 
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
Need a SID admin

```http
DELETE /apiuser/&{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to fetch |

**Code 200 (OK)**
*Content example*
```json
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
```json 
{
  "msg":"success",
}
```
#### Count the number of likes/dislike in a subject
Don't need the SID

```http
GET /api/subject/:id/count
```
**Code : 200 (OK)**

*Content example for a subject with only one UpVote and no DownVote*
```json
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
```json
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

```json
{
    "downvote": false,
    "upvote": true
}
```
#### Search subject
Don't need the SID

Search a subject who contains the word 
```http
GET /api/subject/search/:word
```
**Code : 200 (OK)**
*Content example if we try with the word pirate*

```json

  [{
    "id": "19",
    "title": "PIRATE",
    "description": "JOHNNY",
    "nsfw": 1,
    "image": Return the image in base64,
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
    "image":Return the image in base64,
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
Don't need SID

Get a random subject 
```http
GET /api/subject/GetNBSubject/:nb 
```

**Code : 200 (OK)**
*Content example if we try with 2 as nb*

```json
[{
    "id": "22",
    "title": "La famille pirates",
    "description": "un tres bon dessin animé",
    "nsfw": 0,
    "image": Return the image in base64,
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
    "image": Return the image in base64,
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
Don't need SID 
```http
GET /api/subject/GetLastSubjectUpdate/:nb
```

**Code : 200(OK)**
*Example of content you can get*

```json
[{
    "id": "24",
    "title": "Trackmania",
    "description": "Vroum Vroum la voiture",
    "nsfw": 1,
    "image": Return the image in base64,
    "tags": "#Pilote #Voiture #F1 #Vitesse",
    "upvote " : "",
}, {{
    "id": "23",
    "title": "Foot 2 Rue",
    "description": "Quand t aime le foot et que tu viens de la rue",
    "nsfw": 0,
    "image": return the image in base64,
    "tags": "#Style #Foot #Rue #93",
    "upvotes": "",
    "downvotes": "",
    "publishdate": "11/6/2022",
    "lastpostdate": "2022-06-11 12:39:24",
    "allposts": "#20",
    "owner": "408b95ce-adfe-45e8-b657-b66ed81d76ec"
    }]
```

#### Get subject by User 
Don't need SID 
```http
GET /api/subject/GetSubjectsByUser/:id
```

**Code : 200(OK)**
*Example of content you can get*

```json
[{
    "id": "24",
    "title": "Trackmania",
    "description": "Vroum Vroum la voiture",
    "nsfw": 1,
    "image": Return the image in base64,
    "tags": "#Pilote #Voiture #F1 #Vitesse",
    "upvote " : "",
}]
```

#### Get subject by Tags
Don't need SID 
```http
GET /api/subject/GetSubjectsByTag/:tag
```

**Code : 200(OK)**
*Example of content you can get*

```json
[{
    "id": "29",
    "title": "Titre",
    "description": "desc",
    "nsfw": 1,
    "image": Return the image in base64,
    "tags": "#le#chat#luca",
    "upvote " : "",
}]
```

#### Get all Subjects

Don't need the SID

Show all subject.
```http
  GET /api/subjects
```

**Code : 200 (OK)**
Return all Subject 

#### Create Subject 
Don't need the SID

Create a subject.

```http
  Post /api/subject
```
| Parameter       | Type     | Description                                   |
| :--------       | :------- | :--------------------------------             |
| `title`         | `string` | **Required**. Title of the subject            |
| `Description`   | `string` | **Required**. Description of the subject      |
| `imagedata`     | `string` | **Not Required**. Image of the subject        |
| `nsfw`          | `int`    | **Required**. Is the subject nsfw or not      |
| `tags`          | `string` | **Not Required**. Tags of the subject         |


#### Get subject

Show a subject.

Don't need the SID

```http
  GET /api/subject/${id}
```

| Parameter | Type     | Description                          |
| :-------- | :------- | :--------------------------------    |
| `id`      | `string` | **Required**. Id of subject to fetch |

#### Delete Subject
Need a SID admin

Delete a subject.

```http
  Delete /api/subject/${id}
```

| Parameter | Type     | Description                          |
| :-------- | :------- | :--------------------------------    |
| `id`      | `string` | **Required**. Id of subject to fetch |


**Code : 200 (OK)**

*Content example*
```json
{
    "msg": "success",
}
``` 

#### PUT Subject

Modify a subject


```http
  put /api/subject/${id}
```

| Parameter | Type     | Description                          |
| :-------- | :------- | :--------------------------------    |
| `id`      | `string` | **Required**. Id of subject to fetch |

**Code : 200 (OK)**

*Content example*
```json
{
    "msg": "success",
}
``` 



### Posts

#### Create post 

Need the SID

Create a post.

```http
  Post /api/post/
```

| Parameter     | Type     | Description                                   |
| :--------     | :------- | :--------------------------------             |
| `title`       | `string` | **Required**. Title of the post               |
| `Description` | `string` | **Required**. Description of the post         |
| `nsfw`        | `int`    | **Required**. Is the post nsfw or not         |
| `tags`        | `string` | **Not Required**. Tags of the post            |
| `parent`      | `string` | **Required**. Parent of the post              |
| `image`       | `string` | **Not Required**. Image of the post           |

#### Get all posts
Don't need the SID

Get all post.

```http
  GET /api/posts
```

**Code : 200 (OK)**

*Content example*
```json
{
    "msg": "success",
}
```

#### Get post
Don't need the SID

Show a post.

```http
  GET /api/post/${id}
```

| Parameter | Type     | Description                          |
| :-------- | :------- | :--------------------------------    |
| `id`      | `string` | **Required**. Id of subject to fetch |


#### Delete post
Need a SID admin

Delete a post.

```http
  Delete /api/post/${id}
```

| Parameter | Type     | Description                          |
| :-------- | :------- | :--------------------------------    |
| `id`      | `string` | **Required**. Id of subject to fetch |


#### PUT post
Modify a post
```http
  put /api/post/${id}
```

| Parameter | Type     | Description                          |
| :-------- | :------- | :--------------------------------    |
| `id`      | `string` | **Required**. Id of subject to fetch |

#### upvote
Need the SID

```http
GET /api/post/:id/upvote
```
**Code : 200 (OK)**

*Content example upvote*
```json
{
    "msg": "success",
}
```

#### DownVote
Need the SID

```http
GET /api/post/:id/downvote
```
**Code : 200 (OK)**

*Content example DownVote*
```json
{
    "msg": "success",
}
```

#### Count the number of Up and down vote
Don't need the SID

```http
GET /api/post/:id/count
```
**Code : 200 (OK)**

*Content example for a post with only one UpVote and no DownVote*
```json
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
```json
{
    "downvote": false,
    "upvote": true
}
```
#### Search Post 
Don't need the SID

```http
GET /api/post/search/:word
```
**Code : 200(OK)**

#### Get a random Post
Don't need the SID

Get a random Post 
```http
GET /api/Post/GetNBPost/:nb 
```
**Code : 200(OK)**

#### Get last post

Don't need the SID

Get the last post

```http
GET /api/post/GetLastPost/:nb 
```

**Code : 200(OK)**

#### GetPostBySubject

Don't need the SID 

Get a post by using a subject id

```http
GET /api/post/GetPostBySubject/:id
```
**Code : 200(OK)**

#### Get post by user 
Don't need the SID

Get a post using a user id
```http
GET /api/post/GetPostByUser/:id
```

#### Get post by Tags
Don't need SID 
```http
GET /api/subject/GetPostsByTag/:tag
```

**Code : 200(OK)**
*Example of content you can get*

```json
[{
    "id": "39",
    "title": "title",
    "description": "desc",
    "nsfw": 1,
    "image": Return the image in base64,
    "tags": "#Tesla",
    "upvotes" : "",
    "downvotes" : "",
    "publishdate" : "2022-06-15 19:44:04",
    "comments" : "",
    "parent" : "19",
}]
```

### Comments
Don't need the SID

Show all comment.
```http
  GET /api/comments
```

**Code : 200 (OK)**

Return all Comments


#### Create comment 

Need the SID

Create a comment.

```http
  Post /api/comment/
```

| Parameter  | Type     | Description                           |
| :--------  | :------- | :--------------------------------     |
| `content`  | `string` | **Required**. Content of the comment      |

####  Get comment by id 
Don't need the SID

Show the comment with the given id.
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
```json
{
  "msg":"success",
}
```
#### Dislike a comment 
Need the SID

```http 
GET /api/comment/:id/downvote
```
**Code : 200(OK)**

*Content example*
```json
{
  "msg":"success",
}
```

#### count like/dislike on a comment 
Don't need the SID

```http 
GET /api/comment/:id/count 
```
**Code : 200(OK)**

*Content example for a comment with only one UpVote and no DownVote*
```json 
{
    "UpVote": 1,
    "DownVote": 0,
} 

```

#### Vote on a Comment 
Need the SID

```http
GET /api/comment/:id/vote
```
**Code : 200 (OK)**

*Content example for a comment with only one UpVote and no DownVote*
```json
{
    "downvote": false,
    "upvote": true,
}
```

#### Get comment by post id
Need the SID

Get comments using the id of a post

```http
GET /api/comment/GetCommentByPost/:id
```

**Code : 200(OK)**

*Content example for a post with 2 comment ( here the post with the id 17)*
```json
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
Don't need the SID
 
```http
GET /api/count
```

**Code : 200(OK)**

*Example of content you can get*

```json
{
    "Session": 2,
    "Subject": 7,
    "Post": 20,
    "User": 16,
}
```

#### Count the number of User
Don't need the SID

```http
GET /api/count/user
```

**Code : 200(OK)**

*Example of content you can get*

```json
{
    "Nb": 16,
}
```

#### Count the number of Post
Don't need the SID

```http
GET /api/count/post
```

**Code : 200(OK)**

*Example of content you can get*
```json
{
    "Nb": 20,
}
```

#### Count the number of subject
Don't need the SID

```http
GET /api/count/subject
```

**Code : 200(OK)**

*Example of content you can get*


```json
{
    "Nb": 7,

}
```
#### Count the number of Session ( User connected)
Don't need the SID

```http
GET /api/count/session
```

**Code : 200(OK)**

*Example of content you can get*

```json
{
    "Nb": 2,
}
```