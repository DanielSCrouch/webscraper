<h1 align="center">Vidsy Interview Task</h1>

<p align="center">
  Task to be carried out within back-end engineering interviews 👍
</p>

## Introduction

The Vidsy team would like you to implement a **web crawler** in Go. This is a
classic engineering problem that they feel covers a number of software
engineering principles whilst also being interesting for you to work on!

> A [_web crawler_](https://en.wikipedia.org/wiki/Web_crawler) is the portion
> of a search engine that scans web pages looking for links and then follows
> them. It would normally store the data it finds into some datastore where it
> can be useful.

The homework exercise described below tasks you with making the foundation of a
web crawler. The aim of the exercise is to familiarise yourself with the problem
and codebase, before the second stage of the interview process.

As part of the second stage you will be asked to present your work, which will
followed by a short discussion with the Vidsy team about your approach. More
advanced functionality will then be implemented as part of a collaborative
exercise with the Vidsy team.

The test is an opportunity to showcase your technical ability and highlight how
you approach writing and testing services in Go.

**Treat this like writing production code, and show us your best!**

## Setup

1. Clone this repository locally
1. Create a new private GitHub repository
1. As you complete the task, push changes to your new private GitHub repository

Now you're ready to start the task! Please complete the tasks detailed below.

When complete, add `stevenjack` and `revett` as contributors on the GitHub
repository and email your Vidsy hiring manager to confirm your task is ready to
be reviewed 🎉

## Tasks

1. Parse all the `href` attributes within anchor elements in the retrieved page
1. Output a list of all of the `href` attributes found
1. Refactor the existing code in to a simple CLI which can take a URL to parse
   as an argument

**Note** that the use of the standard library (including `golang.org/x/...`
packages) is strongly encouraged. 3rd-party community packages should be
avoided if possible (apart from testing related packages).

## Existing Code

Vidsy engineers have started off the project with code that fetches the page
source of a webpage.

You can run the code using:

```bash
make run
```

## Note

This problem has already been solved by the Vidsy engineering team and will
never be used in production.
