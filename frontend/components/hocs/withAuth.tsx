import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { useAuth } from '../../hooks/useAuth';

// HOC（高階コンポーネント）を使用して、認証されたユーザーのみがアクセスできるようにする
const withAuth = <P extends object>(WrappedComponent: React.ComponentType<P>): React.FC<P> => {
    // 認証を行うコンポーネント
    const ComponentWithAuth = (props: P) => {
        const [isClient, setIsClient] = useState(false); // クライアントサイドかどうかのフラグ
        const [loading, setLoading] = useState(true); // ローディング状態のフラグ
        const router = useRouter(); // ルーターを取得
        const auth = useAuth(); // カスタムフックで認証情報を取得

        // クライアントサイドであることを設定
        useEffect(() => {
            setIsClient(true);
        }, []);

        // 認証状態を監視し、認証されていない場合はサインインページにリダイレクト
        useEffect(() => {
            if (isClient && !auth.isAuthenticated) {
                router.push('/auth/sign_in');
            } else {
                setLoading(false); // 認証されている場合はローディングを解除
            }
        }, [auth.isAuthenticated, router, isClient]);

        // ローディング中はスピナーを表示
        if (loading) {
            return (
                <div className="flex items-center justify-center min-h-screen">
                    <div className="animate-spin rounded-full h-32 w-32 border-t-2 border-b-2 border-blue-500"></div>
                </div>
            );
        }

        // 認証済みの場合は元のコンポーネントをレンダリング
        return <WrappedComponent {...props} />;
    };

    return ComponentWithAuth;
};

export default withAuth;
