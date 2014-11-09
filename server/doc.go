/*
Palantír Web Service


## API Namespace

        /api/v1


## API Endpoint

- Get a notification:

        GET /:subject

        Status 200 OK
        {
                "title": "message title",
                "content": "message content",
                "created_at": "2014-11-09T15:10:36.784725426+08:00"
        }

        Or

        Status 404 Not Found
        {
                "error": "404 Not Found"
        }


- Create a notification:

        POST /:subject

        Payload:
        {
                "title": "new message title",
                "content": "new message content"
        }

        Status 201 Created
        {
                "id": "1001"
        }



### Response Types

- Message:

        {
                "title": "message title",
                "content": "message content",
                "created_at": "2014-11-09T15:10:36.784725426+08:00"
        }


- Ticket:

        {
                "id": "1001"
        }


- Error:

        {
                "error": "error reason"
        }
*/
package main
