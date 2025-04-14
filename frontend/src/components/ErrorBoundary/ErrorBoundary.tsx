import { Component, ReactNode } from 'react';
import { toast } from 'sonner';

type Props = {
  children: ReactNode;
};

type State = {
  error: Error | null;
};

export class ErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      error: null,
    };
  }

  static getDerivedStateFromError(error: Error): State {
    return { error };
  }

  componentDidCatch(error: Error) {
    toast.error(`An error occurred: ${error.message}`);
  }

  render() {
    const { error } = this.state;
    const { children } = this.props;

    if (error) {
      return (
        <div>
          <h2>Something went wrong.</h2>
          <pre>{error.message}</pre>
        </div>
      );
    }

    return children;
  }
}
