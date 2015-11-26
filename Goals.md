# Goals
* One of our goals is to offload as much work as possible onto other projects. One example is using something like Bootstrap/MaterializeCSS instead of writing or own. This is not only good for our mental health, but it also brings attention to other excellent projects.
* Simplicity
  * We want this to be very easy and intuitive for users
  * In our experience, these systems tend to be difficult for users to navigate
* Cross-platform
  * We want to make sure that this project runs well for all users, whether you are using a browser with NoScript or the newest version of Chrome
  * We would like to have native applications for all major mobile devices (Android, iOS, Windows Phone, etc.) but the website should also be mobile-friendly

# TODO
- [x] Login system
  - A basic login system is now setup.
  - User registration is also possible.
    - We should create a utility/library that makes it easy to register many users at once
- [x] User-classes e.g. Student, Teacher, Teaching-Assistant
  - There are now three user-classes: Administrator, Secretary, and User.
    - Administrator: Has complete control over the system. There should only be a few Administrators
    - Secretary: Has privileges to update user status e.g. status in classes, etc.
    - User: The majority of users fall into this class. Can use the system but not modify it. Students, Teaching-Assistants, and Teachers would be here
  - We will need to allow users to be members of multiple classes, as it is common for someone to be both a student and a teaching-assistant.
  - It may be best to have this run on a per-class basis.
- [ ] Classes
- [ ] Database Stuff
  - We need to decide on a database first. I'm very ignorant about databases, but I hear about SQL a lot, so perhaps using MariaDB 10.x would be best?
  - I've also heard a lot about using Redis(?) in conjunction with a SQL database, so we should investigate this.
  - We could also use a NoSQL database, which seems to be very popular right now.
  - BoltDB is very nice, and is pure Go. However, I worry that it may have poor performance.
- [ ] Views/Frontend
  - One of our goals is to minimize client-side rendering, as this is slow and very bad.
  - I have heard good things about Elm recently, so maybe this should be investigated more.
  - Go's html/template works, but it is basically pure HTML and thus can be very annoying.
  - [Ace](https://github.com/yosssi/ace) is basically Go's version of Jade and seems okay, so perhaps we could use this
