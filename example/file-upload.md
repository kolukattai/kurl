---
# do not change refID, this key is used to connect this api with it's saved response
refID: 3a66f6bd-30e2-4522-8d81-a9e1415f84f4

method: "POST"
url: "{{BASE}}/api"
headers: {
	"Content-Type": "multipart/form-data"
}
formData: [
	{
		key: "file",
		type: "file",
		file: ""
	}
]
---

# file upload
api document goes here