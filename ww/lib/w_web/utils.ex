defmodule WWeb.Utils do
  require Integer

  @doc """
  Count the total pages for a number of total records and records per page.

  ## Examples

      iex> count_page(5, 1)
      5

      iex> count_page(5, 2)
      3

      iex> count_page(5, 3)
      2

  """
  def count_page(records, records_per_page) do
    div(records + records_per_page - 1, records_per_page)
  end

  @doc """
  Generate a list of pagination for given num_of_pages, current_page and display_width,
  display_width is a integer of a tuple of left width and right width.

  ## Examples
      iex> WWeb.Utils.generate_pagination(7, 4, {1, 1})
      [1, "...", 3, 4, 5, "...", 7]

      iex> WWeb.Utils.generate_pagination(9, 5, 3)
      [1, "...", 4, 5, 6, "...", 9]

      iex> WWeb.Utils.generate_pagination(9, 5, 4)
      [1, "...", 3, 4, 5, 6, "...", 9]

  """
  def generate_pagination(num_of_pages, current_page, display_width)
      when 1 <= current_page and current_page <= num_of_pages do
    {n, x, d} = {num_of_pages, current_page, display_width}

    {l, r} =
      case d do
        {l, r} when l >= 0 and r >= 0 ->
          {l, r}

        d when d >= 1 ->
          cond do
            Integer.is_odd(d) ->
              {div(d, 2), div(d, 2)}

            Integer.is_even(d) and x - 1 >= n - x ->
              {div(d, 2), div(d, 2) - 1}

            true ->
              {div(d, 2) - 1, div(d, 2)}
          end
      end

    {ll, rr} =
      cond do
        x - 1 <= n - x ->
          ll = min(l, max(x - 2, 0))
          rr = min(l + r - ll, max(n - x - 1, 0))
          {ll, rr}

        true ->
          rr = min(r, max(n - x - 1, 0))
          ll = min(l + r - rr, max(x - 2, 0))
          {ll, rr}
      end

    l_endpoint = if x > 1, do: [1], else: []
    r_endpoint = if x < n, do: [n], else: []
    l_dot = if x - 2 > ll, do: ["..."], else: []
    r_dot = if n - x - 1 > rr, do: ["..."], else: []
    # compatible for elixir 1.9.0,
    # which only support first..last, not support first..last//step
    a = x - ll
    b = x - 1
    l_expand = if a <= b, do: Enum.to_list(a..b), else: []
    c = x + 1
    d = x + rr
    r_expand = if c <= d, do: Enum.to_list(c..d), else: []

    l_endpoint ++ l_dot ++ l_expand ++ [x] ++ r_expand ++ r_dot ++ r_endpoint
  end

  @doc """
  Get the list of tag names.

  ## Examples

      iex> parse_tags("a")
      {:ok, ["a"]}

      iex> parse_tags("a;b;c")
      {:ok, ["a", "b", "c"]}

      iex> parse_tags(nil)
      {:error, "wrong param type"}

      iex> parse_tags(3)
      {:error, "wrong param type"}

  """
  def parse_tags(tag_names) do
    cond do
      is_binary(tag_names) ->
        {:ok, String.split(tag_names, ";", trim: true)}

      true ->
        {:error, "wrong param type"}
    end
  end

  @doc """
  Get the first and last second for a yearmonth string in format "YYYY-MM".

  Raises `MatchError` if yearmonth is not in corrected format.

  ## Examples

      iex> parse_yearmonth(nil)
      {:error, "nil param yearmonth"}

      iex> parse_yearmonth("2021-11")
      {:ok, ~N[2021-11-01 00:00:00], ~N[2021-11-30 23:59:59]}

      iex> parse_yearmonth("wrong format")
      ** (MatchError)

  """
  def parse_yearmonth(yearmonth) do
    # reference: https://stackoverflow.com/questions/40817583/is-there-a-way-to-match-the-first-n-characters-in-an-elixir-string
    if yearmonth == nil do
      {:error, "nil param yearmonth"}
    else
      <<year::binary-size(4)>> <> "-" <> month = yearmonth
      y = String.to_integer(year)
      m = String.to_integer(month)
      days_of_month = :calendar.last_day_of_the_month(y, m)
      {:ok, start_date} = NaiveDateTime.new(y, m, 1, 0, 0, 0)
      {:ok, end_date_t} = NaiveDateTime.new(y, m, days_of_month, 23, 59, 59)
      # {:ok, start_date, NaiveDateTime.add(end_date_t, 1)}
      {:ok, start_date, end_date_t}
    end
  end
end
