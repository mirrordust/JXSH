defmodule WWeb.Router do
  use WWeb, :router

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_live_flash
    plug :put_root_layout, {WWeb.LayoutView, :root}
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  pipeline :restricted_api do
    plug :check_request_header_exist, header: "authorization"
    plug WWeb.Plugs.AuthenticateUser
  end

  scope "/", WWeb do
    pipe_through :browser

    get "/", PageController, :index
    get "/about", PageController, :about
    get "/:post_view_name", PageController, :show
  end

  scope "/api", WWeb do
    pipe_through :api

    scope "/auth", Auth, as: :auth do
      # resources "/users", UserController, except: [:edit, :new]
      resources "/sessions", SessionController, only: [:create, :delete], singleton: true
    end

    scope "/cms", CMS, as: :cms do
      pipe_through :restricted_api

      resources "/posts", PostController, except: [:edit, :new]
      resources "/tags", TagController, except: [:edit, :new]
    end
  end

  # Enables LiveDashboard only for development
  #
  # If you want to use the LiveDashboard in production, you should put
  # it behind authentication and allow only admins to access it.
  # If your application does not have an admins-only section yet,
  # you can use Plug.BasicAuth to set up some basic authentication
  # as long as you are also using SSL (which you should anyway).
  if Mix.env() in [:dev, :test] do
    import Phoenix.LiveDashboard.Router

    scope "/" do
      pipe_through :browser
      live_dashboard "/dashboard", metrics: WWeb.Telemetry, ecto_repos: [W.Repo]
    end
  end

  defp check_request_header_exist(conn, header: header) do
    # Request fields: https://hexdocs.pm/plug/Plug.Conn.html#module-request-fields
    case get_req_header(conn, header) do
      [] ->
        conn
        |> Phoenix.Controller.json(%{errors: "missing essential header"})
        |> halt()

      _ ->
        conn
    end
  end
end
