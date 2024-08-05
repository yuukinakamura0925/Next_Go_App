'use client';

import React, { useState, useEffect } from 'react';
import apiClient from '../api/client'; // APIクライアントをインポート
import withAuth from '../../components/hocs/withAuth';
import PageTitle from '../../components/atoms/PageTitle';

interface MenuCategory {
  id: number;
  name: string;
}

const MenuCategoryPage: React.FC = () => {
  const [categories, setCategories] = useState<MenuCategory[]>([]);
  const [newCategoryName, setNewCategoryName] = useState<string>('');
  const [isAdding, setIsAdding] = useState<boolean>(false);
  const [editingCategoryId, setEditingCategoryId] = useState<number | null>(null);
  const [editingCategoryName, setEditingCategoryName] = useState<string>('');

  useEffect(() => {
    // カテゴリーデータを取得する
    const fetchCategories = async () => {
      try {
        const response = await apiClient.get('/api/menu_categories');
        // menu_categoriesが配列であるか確認し、セット
        if (Array.isArray(response.data.menu_categories)) {
          setCategories(response.data.menu_categories);
        } else {
          console.error('Invalid response data:', response.data);
        }
      } catch (error) {
        console.error('Error fetching menu categories:', error);
      }
    };
    fetchCategories();
  }, []);

  const handleAddCategory = async () => {
    try {
      const response = await apiClient.post('/api/menu_categories', { name: newCategoryName });

      // レスポンスから新しく追加されたカテゴリデータを正しく取り出す
      const newCategory = response.data.menu_category; // ここで適切なキーを使用
      setCategories((prevCategories) => [...prevCategories, newCategory]); // 既存のカテゴリに新しいカテゴリを追加

      setNewCategoryName('');
      setIsAdding(false);
    } catch (error) {
      console.error('Error adding category:', error);
    }
  };

  const handleEditCategory = async (id: number, name: string) => {
    try {
      await apiClient.patch(`/api/menu_categories/${id}`, { name });
      setCategories(categories.map(category =>
        category.id === id ? { ...category, name } : category
      ));
      setEditingCategoryId(null); // 編集モード終了
    } catch (error) {
      console.error('Error editing category:', error);
    }
  };

  const handleDeleteCategory = async (id: number) => {
    if (!window.confirm('本当にこのカテゴリを削除しますか？')) {
      return; // ユーザーが削除をキャンセルした場合は処理を中断
    }

    try {
      await apiClient.delete(`/api/menu_categories/${id}`);
      setCategories(categories.filter(category => category.id !== id));
      alert('カテゴリが削除されました。');
    } catch (error) {
      console.error('Error deleting category:', error);
      alert('カテゴリの削除中にエラーが発生しました。');
    }
  };

  return (
    <div className="p-4">
      <PageTitle title="メニューカテゴリ" />
      <div className="bg-white p-6 rounded-lg shadow-md">
        <table className="w-full table-auto">
          <thead>
            <tr>
              <th className="px-4 py-2 text-center">カテゴリ名</th>
              <th className="px-4 py-2 text-center">操作</th>
            </tr>
          </thead>
          <tbody>
            {categories.map((category) => (
              <tr key={category.id} className="border-gray-400">
                <td className="border px-4 py-2 text-center border-gray-400">
                  {editingCategoryId === category.id ? (
                    <input
                      type="text"
                      value={editingCategoryName}
                      onChange={(e) => setEditingCategoryName(e.target.value)}
                      className="border p-2 w-full"
                    />
                  ) : (
                    category.name
                  )}
                </td>
                <td className="border px-4 py-2 text-center border-gray-400">
                  {editingCategoryId === category.id ? (
                    <button
                      className="bg-green-500 hover:bg-green-700 text-white font-bold py-1 px-2 rounded mr-2"
                      onClick={() => handleEditCategory(category.id, editingCategoryName)}
                    >
                      更新
                    </button>
                  ) : (
                    <>
                      <button
                        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded mr-2"
                        onClick={() => {
                          setEditingCategoryId(category.id);
                          setEditingCategoryName(category.name);
                        }}
                      >
                        編集
                      </button>
                      <button
                        className="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-2 rounded"
                        onClick={() => handleDeleteCategory(category.id)}
                      >
                        削除
                      </button>
                    </>
                  )}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
        {isAdding && (
          <div className="mt-4">
            <input
              type="text"
              value={newCategoryName}
              onChange={(e) => setNewCategoryName(e.target.value)}
              placeholder="Enter category name"
              className="border p-2 mr-2"
            />
            <button
              className="bg-green-500 hover:bg-green-700 text-white font-bold py-1 px-2 rounded"
              onClick={handleAddCategory}
            >
              保存
            </button>
          </div>
        )}
        <button
          className="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          onClick={() => setIsAdding(true)}
        >
          追加
        </button>
      </div>
    </div>
  );
}

export default withAuth(MenuCategoryPage);
