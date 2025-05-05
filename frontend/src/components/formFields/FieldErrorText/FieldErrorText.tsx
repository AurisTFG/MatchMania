import styles from './FieldErrorText.module.css';

export type FieldErrorTextProps = {
  messages: string[];
  marginTop?: string | number;
};

export default function FieldErrorText({
  messages,
  marginTop,
}: FieldErrorTextProps) {
  if (messages.length === 0) {
    return null;
  }

  return (
    <div
      className={styles.errorContainer}
      style={marginTop ? { marginTop } : undefined}
    >
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
