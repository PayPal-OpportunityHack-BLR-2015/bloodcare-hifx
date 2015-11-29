package com.hifx.user.bloodcare.DatabaseHandler;

import android.content.ContentValues;
import android.content.Context;
import android.database.Cursor;
import android.database.DatabaseUtils;
import android.database.SQLException;
import android.database.sqlite.SQLiteDatabase;
import android.database.sqlite.SQLiteDatabaseLockedException;
import android.database.sqlite.SQLiteException;
import android.database.sqlite.SQLiteOpenHelper;
import android.database.sqlite.SQLiteTableLockedException;
import android.util.Log;
import com.hifx.user.bloodcare.Constants;

import java.util.HashMap;

public class DatabaseHandler extends SQLiteOpenHelper {
	private static DatabaseHandler mInstance = null;
    private static final int DATABASE_VERSION = 1;


    SQLiteDatabase database;

    public static synchronized DatabaseHandler getInstance(Context ctx) {

        // Use the application context, which will ensure that you
        // don't accidentally leak an Activity's context.
        // See this article for more information: http://bit.ly/6LRzfx
        if (mInstance == null) {
            if (ctx != null) {
                mInstance = new DatabaseHandler(ctx);
            }
        }

        return mInstance;
    }

    public DatabaseHandler(Context applicationcontext) {


        super(applicationcontext, "androidsqlite.db", null, DATABASE_VERSION);

    }

	@Override
	public void onCreate(SQLiteDatabase database) {
		String queryapi = "";

		try {
			queryapi = "CREATE TABLE if not exists  api(code TEXT,page INT,parentCode TEXT,newsdata TEXT,expiry TEXT,offline INT,logTime TEXT, PRIMARY KEY (code, page) ON CONFLICT IGNORE)";

			database.execSQL(queryapi);

		} catch (SQLException e) {
			e.printStackTrace();
		}
	}

	@Override
	public void onUpgrade(SQLiteDatabase database, int version_old,
			int current_version) {
		String queryapi;

		queryapi = "DROP TABLE IF EXISTS api";

		database.execSQL(queryapi);

		onCreate(database);
	}

	public void insertNews(String code, int page, String newsdata,
			String expiry, int offline, String logTime, String parentCode) {
		// database = null;
		try {
			if (database == null || !database.isOpen())
				database = this.getWritableDatabase();

			ContentValues values = new ContentValues();
			values.put("code", code);
			values.put("page", page);
			values.put("newsdata", newsdata);
			values.put("expiry", expiry);
			values.put("offline", offline);
			values.put("logTime", logTime);

			values.put("parentCode", parentCode);

			database.insertWithOnConflict("api", null, values,
					SQLiteDatabase.CONFLICT_REPLACE);

		} catch (SQLException e) {
			e.printStackTrace();

		}

	}

	public void clearCache() {

		if (database == null || !database.isOpen())
			database = this.getWritableDatabase();
		database.delete("api", null, null);

	}


    @Override
    public synchronized void close() throws SQLException {
        if (database != null)
            database.close();
        super.close();
    }
	public void deleteParentcode(String parentCode) {
		if (database == null || !database.isOpen())
			database = this.getWritableDatabase();
		String deleteQuery = "DELETE FROM  api where parentCode='" + parentCode
				+ "'";
		database.execSQL(deleteQuery);
	}

	public HashMap<String, String> getNewsdataItem(String articlecode, int page) {
		final HashMap<String, String> data = new HashMap<String, String>();
        Log.i("database", "databse" + database);

		try {
			if (database == null || !database.isOpen())

				database = this.getReadableDatabase();
		}
        catch (SQLiteDatabaseLockedException e) {
            return data;
        }
        catch (SQLiteTableLockedException e) {
            return data;
        }catch (SQLiteException e) {
			return data;
		}

		String newsdata = null;
		String logTime = null;
		String expiry = null;
		Cursor cursor = null;
		try {
			String selectQuery = "SELECT newsdata, logTime,expiry,offline FROM api where code="
					+ DatabaseUtils.sqlEscapeString(articlecode)
					+ " AND page ='" + page + "';";

			cursor = database.rawQuery(selectQuery, null);

		} catch (SQLException e) {
			e.printStackTrace();
		}

		if (cursor != null) {
			if (cursor.getCount() > 0) {
				if (cursor.moveToFirst()) {
					do {
						newsdata = cursor.getString(cursor
								.getColumnIndex("newsdata"));
						logTime = cursor.getString(cursor
								.getColumnIndex("logTime"));
						expiry = cursor.getString(cursor
								.getColumnIndex("expiry"));

						data.put(Constants.newsdata, newsdata);
						data.put(Constants.logTime, logTime);
						data.put(Constants.expiry, expiry);

					} while (cursor.moveToNext());
				}
			}
		}
		if (cursor != null) {
			cursor.close();
		}
		if (database != null) {
			database.close();
		}
		return data;

	}



}