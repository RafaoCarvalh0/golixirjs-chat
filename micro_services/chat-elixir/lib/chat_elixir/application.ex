defmodule ChatElixir.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      ChatElixirWeb.Telemetry,
      ChatElixir.Repo,
      {DNSCluster, query: Application.get_env(:chat_elixir, :dns_cluster_query) || :ignore},
      {Phoenix.PubSub, name: ChatElixir.PubSub},
      # Start the Finch HTTP client for sending emails
      {Finch, name: ChatElixir.Finch},
      # Start a worker by calling: ChatElixir.Worker.start_link(arg)
      # {ChatElixir.Worker, arg},
      # Start to serve requests, typically the last entry
      ChatElixirWeb.Endpoint
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: ChatElixir.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    ChatElixirWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
