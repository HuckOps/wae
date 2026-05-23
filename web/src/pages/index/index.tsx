import AddServiceModal from '@/components/AddServiceModal';
import Pagination from '@/components/Pagination';
import StatusTag from '@/components/StatusTag';
import { createService, getServices } from '@/services/api';
import { FormattedMessage, useIntl } from '@@/plugin-locale';
import { useCallback, useEffect, useState } from 'react';
import styles from './index.less';

interface ServiceFormData {
  name: string;
  repo: string;
  domain: string;
  cluster: string;
  description: string;
}

export default function HomePage() {
  const intl = useIntl();

  const [total, setTotal] = useState<number>(0);
  const [services, setServices] = useState<API.Service[]>([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(10);

  const fetchServices = useCallback(async () => {
    try {
      const resp = await getServices(currentPage, pageSize);
      if (!resp) {
        return;
      }
      setTotal(resp.data.total || 0);
      setServices(resp.data.items || []);
    } catch (error) {
      console.error('Failed to fetch services');
    }
  }, [currentPage, pageSize]);

  useEffect(() => {
    fetchServices();
  }, [fetchServices]);

  const handleSubmit = async (formData: ServiceFormData) => {
    try {
      const resp = await createService(formData);
      if (resp) {
        setIsModalOpen(false);
        fetchServices();
      }
    } catch (error) {
      console.error('Failed to create service');
    }
  };

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  const handlePageSizeChange = (size: number) => {
    setPageSize(size);
    setCurrentPage(1);
  };

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <div>
          <h1 className={styles.title}>
            <FormattedMessage id="service.title" />
          </h1>
        </div>
        <button
          className={styles.addButton}
          onClick={() => setIsModalOpen(true)}
        >
          + <FormattedMessage id="service.addService" />
        </button>
      </div>

      <div className={styles.stats}>
        <div className={styles.statCard}>
          <div className={styles.statTitle}>
            <FormattedMessage id="service.managedServices" />
          </div>
          <div className={styles.statValue}>{total}</div>
        </div>
        <div className={styles.statCard}>
          <div className={styles.statTitle}>
            <FormattedMessage id="service.last24hoursUpdate" />
          </div>
          <div className={styles.statValue}></div>
        </div>
        <div className={styles.statCard}>
          <div className={styles.statTitle}>
            <FormattedMessage id="service.statusFailedServiceCount" />
          </div>
          <div className={styles.statValue}></div>
        </div>
      </div>

      <div className={styles.tableCard}>
        <div className={styles.toolbar}>
          <input
            className={styles.searchInput}
            placeholder={intl.formatMessage({
              id: 'service.serviceSearchPlaceholder',
            })}
          />
          <button className={styles.filterButton}>
            <FormattedMessage id="service.search" />
          </button>
        </div>

        <div className={styles.tableContainer}>
          <table className={styles.table}>
            <thead>
              <tr>
                {['name', 'repo', 'domain', 'cluster', 'status', 'lastDeploy'].map((head) => (
                  <th key={head} className={styles.th}>
                    <FormattedMessage id={`service.${head}`} />
                  </th>
                ))}
              </tr>
            </thead>
            <tbody>
              {services.map((service) => (
                <tr key={service.id} className={styles.tableRow}>
                  <td className={`${styles.td} ${styles.tdFirst}`}>
                    <div className={styles.userInfo}>{service.name}</div>
                  </td>
                  <td className={styles.td}>{service.repo}</td>
                  <td className={styles.td}>{service.domain}</td>
                  <td className={styles.td}>{service.cluster}</td>
                  <td className={styles.td}>
                    <StatusTag status={service.status} />
                  </td>
                  <td className={styles.td}>{service.last_deploy}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <Pagination
          current={currentPage}
          pageSize={pageSize}
          total={total}
          onChange={handlePageChange}
          onPageSizeChange={handlePageSizeChange}
        />
      </div>

      <AddServiceModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onSubmit={handleSubmit}
      />
    </div>
  );
}