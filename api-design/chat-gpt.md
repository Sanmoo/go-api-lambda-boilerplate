Me:

I am starting the design of a REST API for tracking the following:

* Which books have I already read, and which ones I intend to read in the future.
* Which movies have I already watched, and which ones I intend to watch in the future.
* Which tv series have I already watched, and which ones I intend to watch in the future.
* Which electronic games have I already played, and which ones I intend to play in the future.
* Which non electronic games have I already played, and which ones I intend to try in the future.

For each one of them, basides important information (like author for books and platform for electronic games), I also want to track my personal rating (1 to 10 score). Please gimme a main.tsp file describing my API in TypeSpec language.

ChatGPT:

Here is the TypeSpec file for your REST API. It defines resources for books, movies, TV series, electronic games, and non-electronic games, including their key attributes and operations for CRUD functionality. Let me know if you’d like to refine or expand this design!

```
import "@typespec/rest";

using TypeSpec.Rest;

@service({
  title: "MediaTracker",
  version: "1.0.0",
  description: "A REST API for tracking books, movies, TV series, electronic games, and non-electronic games."
})
namespace MediaTracker {

  @resource("books")
  model Book {
    @key
    id: string;

    title: string;
    author: string;
    genre?: string;
    status: "read" | "to-read";
    rating?: int32; // 1 to 10
  }

  @resource("movies")
  model Movie {
    @key
    id: string;

    title: string;
    director?: string;
    genre?: string;
    status: "watched" | "to-watch";
    rating?: int32; // 1 to 10
  }

  @resource("tv-series")
  model TVSeries {
    @key
    id: string;

    title: string;
    seasons?: int32;
    status: "watched" | "to-watch";
    rating?: int32; // 1 to 10
  }

  @resource("electronic-games")
  model ElectronicGame {
    @key
    id: string;

    title: string;
    platform: string;
    genre?: string;
    status: "played" | "to-play";
    rating?: int32; // 1 to 10
  }

  @resource("non-electronic-games")
  model NonElectronicGame {
    @key
    id: string;

    title: string;
    type?: string;
    status: "played" | "to-try";
    rating?: int32; // 1 to 10
  }

  interface MediaOperations {
    @get()
    list(): Book[] | Movie[] | TVSeries[] | ElectronicGame[] | NonElectronicGame[];

    @post()
    create(data: Book | Movie | TVSeries | ElectronicGame | NonElectronicGame): Book | Movie | TVSeries | ElectronicGame | NonElectronicGame;

    @get resource(): Book | Movie | TVSeries | ElectronicGame | NonElectronicGame;

    @put resource(): Book | Movie | TVSeries | ElectronicGame | NonElectronicGame;

    @delete resource(): void;
  }

  extend Book with MediaOperations {}
  extend Movie with MediaOperations {}
  extend TVSeries with MediaOperations {}
  extend ElectronicGame with MediaOperations {}
  extend NonElectronicGame with MediaOperations {}
}
```
