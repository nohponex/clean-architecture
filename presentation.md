# Clean Architecture and Event-Driven Architecture
## Spafaridis Xenofon,
### Director of Software Architecture Vivante Health


----

# Ports & Adapters
![](https://herbertograca.files.wordpress.com/2017/03/hexagonal-arch-5-traditional2.png?w=415)

https://herbertograca.com/2017/09/14/ports-adapters-architecture/

# The Clean Architecture
![The Clean Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html


----

![](https://herbertograca.files.wordpress.com/2018/11/020-explicit-architecture-svg.png?w=1100)
![](https://herbertograca.files.wordpress.com/2018/11/040-explicit-architecture-svg.png?w=1100)
![](https://herbertograca.files.wordpress.com/2018/11/060-explicit-architecture-svg.png?w=1100)
![](https://herbertograca.files.wordpress.com/2018/11/070-explicit-architecture-svg.png?w=1100)

https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/

----

# SOLID Principles

- **Single responsibility**: *a class should only have a single responsibility, that is, only changes to one part of the software's specification should be able to affect the specification of the class*
- **Open–closed**: *software entities ... should be open for extension, but closed for modification*
- **Liskov substitution**: *objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program*
- **Interface segregation**: *many client-specific interfaces are better than one general-purpose interface*
- **Dependency inversion**: *depend upon abstractions, not concretions*

https://en.wikipedia.org/wiki/SOLID

----

https://martinfowler.com/articles/mocksArentStubs.html

----

# Event-driven architecture (EDA)

![](https://chrisrichardson.net/i/sagas/Create_Order_Saga.png)

[GOTO 2017 • The Many Meanings of Event-Driven Architecture • Martin Fowler](https://www.youtube.com/watch?v=STKCRSUsyP0)

## Why do all this?

- Responsiveness. Since everything happens as soon as possible and nobody is waiting on anyone else, event-driven architecture provides the fastest possible response time.
- Scalability. Since you don’t have to consider what’s happening downstream, you can add service instances to scale. Topic routing and filtering can divide up services quickly and easily – as in command query responsibility segregation.
- Agility. If you want to add another service, you can just have it subscribe to an event and have it generate new events of its own. The existing services don’t know or care that this has happened, so there’s no impact on them.
- Agility again. By using an event mesh you can deploy services wherever you want: cloud, on premises, in a different country, etc. Since the event mesh learns where subscribers are, you can move services around without the other services knowing.

https://solace.com/what-is-event-driven-architecture/



![reactive manifesto](https://www.reactivemanifesto.org/images/reactive-traits.svg)
https://www.reactivemanifesto.org/

----

# Books:

|  |  |
|-----------|:-----------:|
|[Growing Object-Oriented Software, Guided by Tests](https://www.goodreads.com/book/show/4268826-growing-object-oriented-software-guided-by-tests)  | ![](https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1348030542l/4268826.jpg) |
|[Clean Architecture](https://www.goodreads.com/book/show/18043011-clean-architecture)                                                              | ![](https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1471680093l/18043011.jpg) |
|[Patterns of Enterprise Application Architecture](https://www.goodreads.com/book/show/70156.Patterns_of_Enterprise_Application_Architecture)       | ![](https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1440294142l/70156.jpg) |
|[ Domain-Driven Design: Tackling Complexity in the Heart of Software](https://www.goodreads.com/book/show/179133.Domain_Driven_Design)             | ![](https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1287493789l/179133.jpg) |
