import { ComponentType } from 'react';
import { ErrorBoundary } from '../components/ErrorBoundary';

export default function withErrorBoundary<P extends object>(
  WrappedComponent: ComponentType<P>,
) {
  const WithErrorBoundary = (props: P) => {
    return (
      <ErrorBoundary>
        <WrappedComponent {...props} />
      </ErrorBoundary>
    );
  };

  WithErrorBoundary.displayName = `WithErrorBoundary(${WrappedComponent.displayName ?? (WrappedComponent.name || 'Component')})`;

  return WithErrorBoundary;
}
