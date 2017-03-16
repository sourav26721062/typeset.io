# Problem Statement

#### Blogging platform with in-place comments
You are to create a blogging product that supports in-place annotations. Requirements in the form of user-stories follow:
An admin should be able to add blog posts. Blog posts have a unique random identifier, a title and plain text where paragraphs are separated by two new-line characters.
A viewer should be able to view all blog posts (list-mode) starting with first 5 and then the next 5 and so on. This view will not have any comments.
A viewer should be able to click on one of these blogs to view it in full-mode. In full-mode, all past comments on the text are visible next to the text. Also, the viewer is able to comment on a paragraph of text. In essence, the comment is on a paragraph.
Nuances:
The API, when it responds with comments should provide markers to indicate the paragraph in question. The notation for the marker is up to you to choose.

Non-requirements: Please do not think of or address any of these items:
There are no users. No authentication. Everyone could be an admin and a viewer.
The comments are not attached to any user, because there are none. There are just blogs and comments.
You are not required to create the UI to achieve any of the goals. Simple APIs are all we are looking for. These could be HTTP APIs. The HTTP APIs should be able to respond to all the user stories enumerated above.
You are not required to write tests for this. You may, if that is how you work.



## API Documentation

**Endpoints** :

- [GET] : /api/post

	description : to get all post
    
    sample request : http://localhost:8080/api/post
    
    sample response : 
    
    ```
    [{
	"id": "586cfd2b9dcec0c436fe7d58",
	"title": "This is My First Blog post",
	"content": [{
		"id": "586cfd2b9dcec0c436fe7d56",
		"text": "Lorem ipsum dolor sit amet, detraxit referrentur signiferumque te pro, wisi appareat facilisis vim ad, nec cu suas errem putent. Mei ut option eleifend, at virtute tincidunt necessitatibus mea, eu verear molestie duo. Usu meis invenire liberavisse eu. Esse maiorum duo no, paulo audiam quaerendum ex eam.",
		"createdOn": "2017-01-04T19:18:27.76+05:30"
	}, {
		"id": "586cfd2b9dcec0c436fe7d57",
		"text": "Lorem ipsum dolor sit amet, detraxit referrentur signiferumque te pro, wisi appareat facilisis vim ad, nec cu suas errem putent. Mei ut option eleifend, at virtute tincidunt necessitatibus mea, eu verear molestie duo. Usu meis invenire liberavisse eu. Esse maiorum duo no, paulo audiam quaerendum ex eam.",
		"comments": [{
			"id": "586d105b9dcec0dc37b0e7e8",
			"text": "I like it",
			"createdOn": "2017-01-04T20:40:19.465+05:30"
		}],
		"createdOn": "2017-01-04T19:18:27.76+05:30"
	}],
	"createdOn": "2017-01-04T19:18:27.76+05:30",
	"link": "http://localhost:8080/api/post/586cfd2b9dcec0c436fe7d58"}, {
	"id": "586cdb369dcec07eb5d6ba93",
	"title": "This is my second Blog Post",
	"content": [{
		"id": "586cdb369dcec07eb5d6ba92",
		"text": "Lorem ipsum dolor sit amet, detraxit referrentur signiferumque te pro, wisi appareat facilisis vim ad, nec cu suas errem putent. Mei ut option eleifend, at virtute tincidunt necessitatibus mea, eu verear molestie duo. Usu meis invenire liberavisse eu. Esse maiorum duo no, paulo audiam quaerendum ex eam.",
		"comments": [{
			"id": "586d08069dcec0d3bbc46e49",
			"createdOn": "2017-01-04T20:04:46.339+05:30"
		}],
		"createdOn": "2017-01-04T16:53:34.446+05:30"
	}],
	"createdOn": "2017-01-04T16:53:34.446+05:30",
	"link": "http://localhost:8080/api/blog/586cdb369dcec07eb5d6ba93"}]
    ```
 
 
- [GET] : /api/post?page=1

	description : to get all post based on pagination (Each page returns 5 result)
    
    sample request : http://localhost:8080/api/post?page=1
    
    sample response : 
    
    ```
    [{
	"id": "586cfd2b9dcec0c436fe7d58",
	"title": "This is My First Blog post",
	"content": [{
		"id": "586cfd2b9dcec0c436fe7d56",
		"text": "Lorem ipsum dolor sit amet, detraxit referrentur signiferumque te pro, wisi appareat facilisis vim ad, nec cu suas errem putent. Mei ut option eleifend, at virtute tincidunt necessitatibus mea, eu verear molestie duo. Usu meis invenire liberavisse eu. Esse maiorum duo no, paulo audiam quaerendum ex eam.",
		"createdOn": "2017-01-04T19:18:27.76+05:30"
	}, {
		"id": "586cfd2b9dcec0c436fe7d57",
		"text": "Lorem ipsum dolor sit amet, detraxit referrentur signiferumque te pro, wisi appareat facilisis vim ad, nec cu suas errem putent. Mei ut option eleifend, at virtute tincidunt necessitatibus mea, eu verear molestie duo. Usu meis invenire liberavisse eu. Esse maiorum duo no, paulo audiam quaerendum ex eam.",
		"comments": [{
			"id": "586d105b9dcec0dc37b0e7e8",
			"text": "I like it",
			"createdOn": "2017-01-04T20:40:19.465+05:30"
		}],
		"createdOn": "2017-01-04T19:18:27.76+05:30"
	}],
	"createdOn": "2017-01-04T19:18:27.76+05:30",
	"link": "http://localhost:8080/api/post/586cfd2b9dcec0c436fe7d58"}, {
	"id": "586cdb369dcec07eb5d6ba93",
	"title": "This is my second Blog Post",
	"content": [{
		"id": "586cdb369dcec07eb5d6ba92",
		"text": "Lorem ipsum dolor sit amet, detraxit referrentur signiferumque te pro, wisi appareat facilisis vim ad, nec cu suas errem putent. Mei ut option eleifend, at virtute tincidunt necessitatibus mea, eu verear molestie duo. Usu meis invenire liberavisse eu. Esse maiorum duo no, paulo audiam quaerendum ex eam.",
		"comments": [{
			"id": "586d08069dcec0d3bbc46e49",
			"createdOn": "2017-01-04T20:04:46.339+05:30"
		}],
		"createdOn": "2017-01-04T16:53:34.446+05:30"
	}],
	"createdOn": "2017-01-04T16:53:34.446+05:30",
	"link": "http://localhost:8080/api/blog/586cdb369dcec07eb5d6ba93"}]
    ```


- [POST] : /api/post

	description : create a new blog post
    
    sample request : http://localhost:8080/api/post
    
    sample request body :
    
    ```
    {
    "title" : "This is my first blog",
    "content" :"hi/n/nworld"
    }
    ```
    
    sample response : 
    
    ```
    {
    	"status" : "Success"
    }
    
    ```
    
    
- [GET] : /api/post/:id

	description : get a specific blog post
    
    sample request : http://localhost:8080/api/post/586cc1859dcec06edc261764
    
    sample request body :
    
    ```
    {
    "title" : "This is my first blog",
    "content" :"hi/n/nworld"
    }
    ```
    
    sample response : 
    
    ```
    {
  	"status": "Fail",
  	"error": "post doesnot exist"
	}
    
    ```
    
- [POST] : /api/para/:pid/comment

	description : post a comment on a specific paragraph
    
    sample request : http://localhost:8080/api/para/586cc1859dcec06edc261764/comment
    
    sample request body :
    
    ```
    {
    "text" : "my first comment"
    }
    ```
    
    sample response : 
    
    ```
    {
  	"status": "Success"
	}
    
    ```
    
    
- [GET] : /api/para/:pid/comment

	description : get all comments on a specific paragraph
    
    sample request : http://localhost:8080/api/para/586cc1859dcec06edc261764/comment
    
    sample response : 
    
    ```
    [
  {
    "id": "586cdb369dcec07eb5d6ba92",
    "text": "Lorem ipsum dolor sit amet, detraxit referrentur signiferumque te pro, wisi appareat 		facilisis vim ad, nec cu suas errem putent. Mei ut option eleifend, at virtute tincidunt 		necessitatibus mea, eu verear molestie duo. Usu meis invenire liberavisse eu. Esse maiorum duo 		no, paulo audiam quaerendum ex eam",
    "comments": [
      {
        "id": "586d08069dcec0d3bbc46e49",
        "text": "My first comment",
        "createdOn": "2017-01-04T20:04:46.339+05:30"
      }
    ],
    "createdOn": "2017-01-04T16:53:34.446+05:30"
  }]
    ```

