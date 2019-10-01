# Steplog

Steplog is a logging package I made to make hierarchical logging easier.

Every line has a parent, from which it is indented, except for the root.

This makes it so that if multiple successive operations are made, with each of them
having the same log format, it is easy to see which line belongs to which operation,
with each of them having a seperate parent log line.
