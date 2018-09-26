# srvdir

srvdir server the current directory via HTTP.

Example of use:

    % ls
    index.html
    post.html
    style.css
    % srvdir -http=:8080

This will serve the files index.html, post.html and style.css in the port 8080,
so you can point your browser to the url http://localhost:8080/post.html and see
the post.html file contents.

It's similar to Python's SimpleHTTPServer, but faster.
