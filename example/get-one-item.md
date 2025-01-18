---
# do not change refID, this key is used to connect this api with it's saved response
refID: 0193edca-cf3e-4634-9832-7fa5bf1bbcbf

method: GET
url: "http://localhost:3000/api/items/{{id}}"
params: {
  id: 21
}
queryParams: {
  limit: 10,
  skip: 10
}
headers: {
  "Content-Type": "application/json"
}

---

# Example get API
some content

```json
{
  "name": "mate"
}
```

