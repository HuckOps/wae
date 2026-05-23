import { FormattedMessage } from '@@/plugin-locale';
import styles from './Pagination.less';

interface PaginationProps {
  current: number;
  pageSize: number;
  total: number;
  onChange: (page: number) => void;
  onPageSizeChange?: (size: number) => void;
}

const pageSizeOptions = [10, 20, 50, 100];

export default function Pagination({
  current,
  pageSize,
  total,
  onChange,
  onPageSizeChange,
}: PaginationProps) {
  const totalPages = Math.ceil(total / pageSize);

  if (totalPages < 1) {
    return null;
  }

  const pages: (number | string)[] = [];
  const showPages = 5;

  if (totalPages <= showPages + 2) {
    for (let i = 1; i <= totalPages; i++) {
      pages.push(i);
    }
  } else {
    pages.push(1);

    if (current > 3) {
      pages.push('...');
    }

    const start = Math.max(2, current - 1);
    const end = Math.min(totalPages - 1, current + 1);

    for (let i = start; i <= end; i++) {
      pages.push(i);
    }

    if (current < totalPages - 2) {
      pages.push('...');
    }

    pages.push(totalPages);
  }

  return (
    <div className={styles.paginationWrapper}>
      <div className={styles.pageSizeSelector}>
        <span className={styles.pageSizeLabel}>
          <FormattedMessage id="pagination.pageSize" />
        </span>
        <select
          className={styles.pageSizeSelect}
          value={pageSize}
          onChange={(e) => onPageSizeChange?.(Number(e.target.value))}
        >
          {pageSizeOptions.map((size) => (
            <option key={size} value={size}>
              {size}
            </option>
          ))}
        </select>
      </div>

      <div className={styles.pagination}>
        <button
          className={styles.pageBtn}
          disabled={current === 1}
          onClick={() => onChange(current - 1)}
        >
          <FormattedMessage id="pagination.prev" />
        </button>

        {pages.map((page, index) =>
          typeof page === 'number' ? (
            <button
              key={index}
              className={`${styles.pageBtn} ${
                current === page ? styles.active : ''
              }`}
              onClick={() => onChange(page)}
            >
              {page}
            </button>
          ) : (
            <span key={index} className={styles.ellipsis}>
              {page}
            </span>
          ),
        )}

        <button
          className={styles.pageBtn}
          disabled={current === totalPages}
          onClick={() => onChange(current + 1)}
        >
          <FormattedMessage id="pagination.next" />
        </button>
      </div>
    </div>
  );
}
