import "@typespec/rest";
import "@typespec/versioning";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.Versioning;

@service({
  title: "Media Tracker",
})
@versioned(Versions)
@server("https://example.com", "Single server endpoint")
namespace MediaTracker;

// Enums
enum Versions {
  v010: "0.1.0",
}


enum BookStatus {
  read,
  to_read,
  reading,
  paused,
  gave_up,
}


enum MovieStatus {
  watched,
  to_watch,

  @doc("Means you started watching it but you did not find it interesting enough to complete.")
  gave_up_incomplete,
}

enum TVSeriesStatus {
  ...MovieStatus,
  watching,
}

enum ElectronicGameStatus {
  to_try,
  tried,
  to_beat,
  beated,
  platinated,
  gave_up,
}

enum NonElectronicGameStatus {
  to_try,
  tried,
}

enum NonElectronicGameType {
  board,
  card,
  dice,
  other,
}

model Media {
  @visibility(Lifecycle.Read)
  @key
  id: string;

  title: string;

  @minValue(1)
  @maxValue(10)
  rating?: int32;
}

@resource("books")
@example(#{
  author: "Immanuel Kant",
  genre: "Philosophy",
  status: BookStatus.read,
  title: "Critique of Pure Reason",
  id: "123",
  rating: 8,
})
model Book extends Media {
  author: string;
  genre: string;
  status: BookStatus;
}

@resource("movies")
@example(#{
  director: "Christopher Nolan",
  genre: "Action",
  status: MovieStatus.watched,
  title: "Inception",
  id: "123",
  rating: 7,
})
model Movie extends Media {
  director: string;
  genre: string;
  status: MovieStatus;
}

@resource("tv-series")
@example(#{
  seasons: 5,
  finished: true,
  status: TVSeriesStatus.watched,
  title: "Breaking Bad",
  id: "123",
  rating: 7,
})
model TVSeries extends Media {
  seasons: int32;
  finished: boolean;
  status: TVSeriesStatus;
}

@resource("electronic-games")
@example(#{
  platform: "PlayStation 4/5",
  genre: "Action",
  status: ElectronicGameStatus.beated,
  title: "The Last of Us",
  id: "123",
  rating: 7,
})
model ElectronicGame extends Media {
  platform: string;
  genre?: string;
  status: ElectronicGameStatus;
}

@resource("non-electronic-games")
@example(#{
  type: NonElectronicGameType.board,
  status: NonElectronicGameStatus.tried,
  title: "Chess",
  id: "123",
  rating: 10,
})
model NonElectronicGame extends Media {
  type: NonElectronicGameType;
  status: NonElectronicGameStatus;
}

model Ok<T> {
  ...OkResponse;
  ...Body<T>;
}

model Created<T> {
  ...CreatedResponse;
  ...Body<T>;
}

model Deleted<T> {
  ...NoContentResponse;
}

model Updated<T> {
  ...AcceptedResponse;
}

model ResourceCreated<T extends Media> {
  ...CreatedResponse;
  ...T;
}

model ResourceList<T extends Media> {
  ...OkResponse;
  @pageItems items: T[];
}

model ResourceAccepted<T> {
  ...AcceptedResponse;
  ...Body<T>;
}

@error
model ResourceNotFound {
  ...NotFoundResponse;
  ...Body<{
    code: "NOT_FOUND";
    message: string;
  }>;
}

@error
model ResourceValidationError {
  ...BadRequestResponse;
  ...Body<{
    code: "VALIDATION_ERROR";
    message: string;
    details: string[];
  }>;
}

@error
model UnauthorizedErrorResponse {
  ...UnauthorizedResponse;
  code: "UNAUTHORIZED";
  message: string;
}

@error
model InternalServerErrorResponse {
  @statusCode statusCode: 500;
  ...Body<{
    code: "INTERNAL_SERVER_ERROR";
    message: string;
  }>;
}

@useAuth(BearerAuth)
interface Crud<T extends Media> {
  @post
  create(@body data: T):
    | ResourceCreated<T>
    | ResourceValidationError
    | InternalServerErrorResponse
    | UnauthorizedErrorResponse;

  @get
  read(@path id: string):
    | Ok<T>
    | ResourceNotFound
    | InternalServerErrorResponse
    | UnauthorizedErrorResponse;

  @put
  put(@path id: string, @body data: T):
    | ResourceAccepted<T>
    | ResourceValidationError
    | InternalServerErrorResponse
    | UnauthorizedErrorResponse;

  @list
  @get
  op list(
    @query @offset skip?: int32 = 0,
    @query @pageSize perPage?: int32 = 100,
  ): ResourceList<T> | InternalServerErrorResponse | UnauthorizedErrorResponse;

  @delete
  op delete(
    @path id: string,
  ): Deleted<T> | InternalServerErrorResponse | UnauthorizedErrorResponse;
}

@route("/books")
@tag("Books")
interface Books extends Crud<Book> {}

@route("/electronic-games")
@tag("Electronic Games")
interface ElectronicGames extends Crud<ElectronicGame> {}

@route("/non-electronic-games")
@tag("Non-Electronic Games")
interface NonElectronicGames extends Crud<NonElectronicGame> {}

@route("/movies")
@tag("Movies")
interface Movies extends Crud<Movie> {}

@route("/tv-series")
@tag("TV Series")
interface TV_Series extends Crud<TVSeries> {}
