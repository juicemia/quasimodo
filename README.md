quasimodo
=========
a deployment tool for hugo sites on s3

I got tired of deleting all my files using that horrid AWS web console
and then dragging and dropping all my files back in, so I decided I would
write a little go tool to make my uploads easy.

## Requirements

1. `hugo` must be in your `$PATH`. `quasimodo` uses the hugo command to generate your site.

1. `quasimodo` needs its own user with access to your S3 bucket. You can create a new user
for this purpose using Identity and Access Management (IAM) in the AWS console. All he needs
is full permissions on S3 (policy name is AmazonS3FullAccess).

1. Set up your `~/.aws/credentials` file with a profile for `quasimodo`. It should look something
like: 

```
[quasimodo]
aws_access_key_id = <QUASIMODO'S ACCESS KEY>
aws_secret_access_key = <QUASIMODO'S SECRET KEY>
```

With the above, just navigate to your Hugo site's project directory (not public) and run `quasimodo --bucket example.com`. He'll build your site for you and push it to S3.
