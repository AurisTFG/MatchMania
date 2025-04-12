import { render, screen } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";
import { useAuth } from "../../providers/AuthProvider";
import HomePage from "./HomePage";

vi.mock("../../providers/AuthProvider", () => ({
  useAuth: vi.fn(),
}));

describe("HomePage Component", () => {
  afterEach(() => {
    vi.resetAllMocks();
  });

  it("should render welcome message with username when user is logged in", () => {
    (useAuth as unknown as ReturnType<typeof vi.fn>).mockReturnValue({
      user: { username: "TestUser" },
      isLoggedIn: true,
      isLoading: false,
    });

    render(
      <BrowserRouter>
        <HomePage />
      </BrowserRouter>,
    );

    expect(screen.getByText(/Welcome to MatchMania/i)).toBeInTheDocument();
    expect(screen.getByText(/TestUser!/i)).toBeInTheDocument();
    expect(screen.queryByText("Login")).not.toBeInTheDocument();
    expect(screen.queryByText("Sign Up")).not.toBeInTheDocument();
  });

  it('should render welcome message with "Guest" when user is not logged in', () => {
    (useAuth as unknown as ReturnType<typeof vi.fn>).mockReturnValue({
      user: null,
      isLoggedIn: false,
      isLoading: false,
    });

    render(
      <BrowserRouter>
        <HomePage />
      </BrowserRouter>,
    );

    expect(screen.getByText(/Welcome to MatchMania/i)).toBeInTheDocument();
    expect(screen.getByText(/Guest!/i)).toBeInTheDocument();
    expect(screen.getByText("Login")).toBeInTheDocument();
    expect(screen.getByText("Sign Up")).toBeInTheDocument();
  });
});
