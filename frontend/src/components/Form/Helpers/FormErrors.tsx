import styles from './FormErrors.module.css';

export default function FormErrors({ messages }: { messages: string[] }) {
  if (messages.length === 0) {
    return null;
  }

  return (
    <div className={styles.errorContainer}>
      {messages.map((message, index) => (
        <div
          key={index}
          className={styles.errorMessage}
        >
          {message}
        </div>
      ))}
    </div>
  );
}
