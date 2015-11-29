package com.hifx.user.bloodcare.Routers;

import android.content.Context;
import android.database.SQLException;
import android.os.AsyncTask;
import android.os.SystemClock;
import android.util.Log;

import com.hifx.user.bloodcare.Constants;
import com.hifx.user.bloodcare.DatabaseHandler.DatabaseHandler;

import java.util.HashMap;

class GetDatafromDB extends
        AsyncTask<HashMap<String, Object>, Integer, HashMap<String, Object>> {
	private Context applicationContext;
	private AsyncResponseInterface asyncInterface;

	public GetDatafromDB(Context context, AsyncResponseInterface callback) {
        Log.e("application in ", "" + context);
		applicationContext = context;
		asyncInterface=callback;

	}

	@Override
	protected HashMap<String, Object> doInBackground(
			HashMap<String, Object>... params) {
		HashMap<String, String> data = null;
		HashMap<String, Object> arrayListVal = params[0];
		String code = (String) arrayListVal.get(Constants.code);
       Log.i("url from getdata", arrayListVal.get(Constants.url).toString());
		int pagenumber = Integer.valueOf((String) arrayListVal
                .get(Constants.pagenumber));
        if(applicationContext==null){
             return null;
        }
		DatabaseHandler dbhandler = new DatabaseHandler(applicationContext);
		if (dbhandler != null) {
			try {
				if (code != null) {

					data = dbhandler.getNewsdataItem(code, pagenumber);
				}

			} catch (SQLException e) {
				e.printStackTrace();
			}
		}

		if (data != null) {
			if (data.size() > 0) {

				arrayListVal.put(Constants.dbdata, "DBDATA");
				arrayListVal.put(Constants.newsdata,
						data.get(Constants.newsdata));
				
				arrayListVal.put(Constants.expiry, data.get(Constants.expiry));
				arrayListVal
						.put(Constants.logTime, data.get(Constants.logTime));

			} else {
				arrayListVal.put(Constants.dbdata, "NODATA");
			}
		} else {
			arrayListVal.put(Constants.dbdata, "NODATA");

		}
		Log.e("dataaa", "data" + arrayListVal);

		data = null;
		return arrayListVal;
	}

	@Override
	protected void onPreExecute() {
		super.onPreExecute();

	}

	@Override
	protected void onPostExecute(HashMap<String, Object> result) {
		

		
		if(result==null){
			asyncInterface.showToast();
			return;
		}

		if (((String) result.get(Constants.dbdata)).equals("NODATA")) {
			Log.e("NODATA", "parentCode" + result.get(Constants.parentCode));
			Log.e("NODATA", "url" + result.get(Constants.url));
			Log.e("NODATA", "pagenumber" + result.get(Constants.pagenumber));
			Log.e("NODATA", "code" + result.get(Constants.code));
			String parentcode = "";
			if (result.get(Constants.parentCode) != null) {
				parentcode = result.get(Constants.parentCode).toString();
			}
			asyncInterface.getJson((String) result.get(Constants.url),

			(String) result.get(Constants.code),
					(String) result.get(Constants.pagenumber), true,
					parentcode,(String) result.get(Constants.classfunctionname), "retain");

		} else if (result.get(Constants.dbdata).equals("DBDATA")) {
			Log.e("dbdata", "dbdata");
			
			long ctime = SystemClock.elapsedRealtime() / 1000;
			if (ctime <= Long.valueOf(result.get(Constants.logTime).toString())) {
				Log.e("dbdata", "ctime<=");

				asyncInterface.getJson((String) result.get(Constants.url),

				(String) result.get(Constants.code), (String) result
						.get(Constants.pagenumber), true,
						result.get(Constants.parentCode).toString(),(String) result.get(Constants.classfunctionname), "retain");

			} else if (ctime > Long.valueOf(result.get(Constants.expiry)
                    .toString())) {
				Log.e("dbdata", "ctime>=");
				asyncInterface.getJson((String) result.get(Constants.url),

				(String) result.get(Constants.code), (String) result
						.get(Constants.pagenumber), true,
						result.get(Constants.parentCode).toString(),(String) result.get(Constants.classfunctionname), "delete");

			} else {
				String newsdata = (String) result.get(Constants.newsdata);
				if (newsdata != null) {
					asyncInterface.dataFromDB(newsdata,(String) result.get(Constants.classfunctionname));
				}
			}
		}
	}

}
