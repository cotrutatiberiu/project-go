    /*"start": "nodemon --require dotenv/config --inspect=127.0.0.1:5112 index.js",*/

for Data Errors
400 requested information is incomplete or malformed.
422 requested information is okay, but invalid.
423 cant provide token
404 everything is okay, but the resource doesn’t exist.
409 conflict of data exists, even with valid information.

for Auth Errors
401 access token isn’t provided, or is invalid.
403 token is valid, but requires more privileges.

for Standard Statuses
200 everything is okay.
201 created
204 everything is okay, but there’s no content to return.
500 the server throws an error, completely unexpected.

//Token template----------------------------------------------------------------
{
  payload
},
key,
{
  expiresIn: process.env.JWT_ACCESS_KEY_EXPIRE,
  algorithm: 'HS256',
  issuer: "access-token-middleware-backend",
  audience: "web-frontend",
  subject: "access-resources",
  header: {
      typ: "JWT_ACCESS",
      csrf: process.env.CSRF_TOKEN,
      roles: undefined
  }
}

error={
name: "JsonWebTokenError"
message: "invalid signature",
}

error= {
  name: 'TokenExpiredError',
  message: 'jwt expired',
  expiredAt: 1408621000
}

//Temporary store----------------------------------------------------------------
request._client = db connection
request._account = account found
request._response = response to be sent

//Session storage----------------------------------------------------------------
c_t - x-csrf-token
u_i - user info

//Cookies----------------------------------------------------------------
a_t - access cookie
r_t - refresh cookies


//category - domain - subdomain----------------------------------------------------------------

constructii
-materiale
-echipamente

submitForm
handleInputChange

opiniatatest@gmail.com

//Postgresql----------------------------------------------------------------

RowCtor: null
command: "INSERT"
fields: []
oid: 0
rowAsArray: false
rowCount: 1
rows: []
_parsers: []
_types: TypeOverrides {_types: {…}, text: {…}, binary: {…}}

RowCtor: null
command: "UPDATE"
fields: []
oid: null
rowAsArray: false
rowCount: 0
rows: []
_parsers: []
_types: TypeOverrides {_types: {…}, text: {…}, binary: {…}}
__proto__: Object

RowCtor: null
command: "SELECT"
fields: Array(11)
0: Field {name: "account_id", tableID: 16479, columnID: 1, dataTypeID: 23, dataTypeSize: 4, …}
1: Field {name: "first_name", tableID: 16479, columnID: 2, dataTypeID: 1043, dataTypeSize: -1, …}
2: Field {name: "last_name", tableID: 16479, columnID: 3, dataTypeID: 1043, dataTypeSize: -1, …}
3: Field {name: "username", tableID: 16479, columnID: 4, dataTypeID: 1043, dataTypeSize: -1, …}
4: Field {name: "email", tableID: 16479, columnID: 5, dataTypeID: 1043, dataTypeSize: -1, …}
5: Field {name: "language", tableID: 16479, columnID: 6, dataTypeID: 1043, dataTypeSize: -1, …}
6: Field {name: "password", tableID: 16479, columnID: 7, dataTypeID: 1043, dataTypeSize: -1, …}
7: Field {name: "active", tableID: 16479, columnID: 8, dataTypeID: 16, dataTypeSize: 1, …}
8: Field {name: "admin_level", tableID: 16479, columnID: 9, dataTypeID: 23, dataTypeSize: 4, …}
9: Field {name: "created", tableID: 16479, columnID: 10, dataTypeID: 1184, dataTypeSize: 8, …}
10: Field {name: "updated", tableID: 16479, columnID: 11, dataTypeID: 1184, dataTypeSize: 8, …}
length: 11
__proto__: Array(0)
oid: null
rowAsArray: false
rowCount: 1
rows: Array(1)
0: {account_id: 11, first_name: "a", last_name: "a", username: "a", email: "mrtornnado@gmail.com", …}
length: 1

//ERROR CODES 
23505 - duplicate code, unique constraint, unique_violation
code: "23505"
column: undefined
constraint: "accounts_username_key"
dataType: undefined
detail: "Key (username)=(a) already exists."
file: "d:\pginstaller.auto\postgres.windows-x64\src\backend\access\nbtree\nbtinsert.c"
hint: undefined
internalPosition: undefined
internalQuery: undefined
length: 278
line: "534"
name: "error"
position: undefined
routine: "_bt_check_unique"
schema: "public"
severity: "ERROR"
table: "accounts"
where: undefined
message: "duplicate key value violates unique constraint "accounts_username_key""
stack: "error: duplicate key value violates unique constraint "accounts_username_key"↵    at Connection.parseE (D:\Work\project\opinia-ta-server\node_modules\pg\lib\connection.js:604:11)↵    at Connection.parseMessage (D:\Work\project\opinia-ta-server\node_modules\pg\lib\connection.js:401:19)↵    at Socket.<anonymous> (D:\Work\project\opinia-ta-server\node_modules\pg\lib\connection.js:121:22)↵    at Socket.emit (events.js:198:13)↵    at addChunk (_stream_readable.js:288:12)↵    at readableAddChunk (_stream_readable.js:269:11)↵    at Socket.Readable.push (_stream_readable.js:224:10)↵    at TCP.onStreamRead [as onread] (internal/stream_base_commons.js:94:17)"


//Token system
1.User logs in
  server creates:
  -access token
  -refresh token
  -stores the refresh token in db

2.User does actions, access token expires
  -server interogates db if refresh token is still valid
  -if yes give new refresh token
  -if not

//React response----------------------------------------------------------------
data: {status: 201, location: "SIGNUP-CONTROLLER_ACCOUNT-CREATED", description: "Account created successfully.", payload: {…}}
status: 201
statusText: "Created"
headers: {content-length: "272", content-type: "application/json; charset=utf-8"}
config: {url: "http://localhost:5111/api/signup", method: "post", data: "{"payload":{"userName":"usernamu","email":"mrtornn…":"Cotruta","lastName":"Andrei","language":"ro"}}", headers: {…}, transformRequest: Array(1), …}
request: XMLHttpRequest {readyState: 4, timeout: 0, withCredentials: false, upload: XMLHttpRequestUpload, onreadystatechange: ƒ, …}