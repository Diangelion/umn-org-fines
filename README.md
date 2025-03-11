# ForFe

## Description

ForFe is a website application that is used for Multimedia Nusantara University organizations to manage their members' fines, inspired by Discord server concept. ForFe is built using plain HTML for the Frontend, with additional libraries, which are: HTMX, Tailwind, and Local Forage. For the server, ForFe has two servers which are the gateway and the backend. LastThis website application features cover all CRUD functioning that will be detailed on the explanations below by each role.

### Debt Collector (a college student that is/are chamberlain)

1. Create organization, fine list
2. Read organization data, such as: fine amount, members, uploaded fine payment
3. Update organization data, such as: name, fine list, members' role, members' fine status
4. Delete organization, fine list, member

### Organization member (a college student)

1. Join organization using invitational code
2. Upload fine payment in organization
3. Read their fine list, status, and amount
4. Leave organization

## Structure

### A. Frontend

1. Built using HTMX, Tailwind, JS, and Local Forage
2. Using http-server-spa to run the web within desired port

### B. Gateway

1. Built using Go and HTML template
2. Act as a middleman between Frontend and true Backend
3. Get all requests from Frontend, validation, then pass to Backend
4. Process JSON response from Backend into HTML document
5. Response to the Frontend in HTML document form (Server Side Rendering)

### C. Backend

1. Get requests from Gateway
2. Process requests without any validation
3. Response to the Gateway in JSON form

## Tributes

1. ChatGPT -> Go setup, bug/problem fixing, learn Go & HTMX
2. Stackoverflow -> bug/problem fixing
3. Github Discussion -> bug/problem fixing
4. https://flowbite.com/icons/ -> as free icon provider
5. https://ckeditor.com/ -> as free text editor

## Color Palette

1. neutral-100 -> #f5f5f5
2. gray-300 -> #d1d5dc
3. gray-400 -> #99a1af
4. gray-700 -> #364153
5. red-700 -> #c10007

## Self Notes

1. Gateway handle error inside itself -> "Error"
2. Gateway handle error from backend -> "Failed"
3. Gateway always create custom error message & log the actual error
4. Backend always return & log the actual error in each function
5. Backend create custom error message for returning to gateway in JSON
6. In both Gateway and Backend, log doesn't need (.) otherwise returned custom error use (.)
7. Utils or anything besides Internal folder doesn't need to log
