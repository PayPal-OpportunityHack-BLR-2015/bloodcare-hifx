/*
 * Copyright (C) 2015 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.hifx.user.bloodcare;

import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v4.app.Fragment;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.ListView;


import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class ListFragment extends Fragment {
    CustomListView customlistView;

    @Nullable
    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
        ListView rv = (ListView) inflater.inflate(
                R.layout.fragment_list, container, false);
        String[] values = new String[] { "Android List View",
                "Adapter implementation",
                "Simple List View In Android",
                "Create List View Android",
                "Android Example",
                "List View Source Code",
                "List View Array Adapter",
                "Android Example List View"
        };

        // Define a new Adapter
        // First parameter - Context
        // Second parameter - Layout for the row
        // Third parameter - ID of the TextView to which the data is written
        // Forth - the Array of data

        ArrayAdapter<String> adapter = new ArrayAdapter<String>(getActivity(),
                android.R.layout.simple_list_item_1, android.R.id.text1, values);


        // Assign adapter to ListView
        rv.setAdapter(adapter);
        View view = inflater.inflate(R.layout.list_fragment, container, false);
        return view;

        //setupRecyclerView(rv);
        return rv;
    }

//    private void setupRecyclerView(RecyclerView recyclerView) {
//        recyclerView.setLayoutManager(new LinearLayoutManager(recyclerView.getContext()));
//        recyclerView.setAdapter(new SimpleStringRecyclerViewAdapter(getActivity(),
//                getRandomSublist(Cheeses.sCheeseStrings, 30)));
//    }

    private List<String> getRandomSublist(String[] array, int amount) {
        ArrayList<String> list = new ArrayList<>(amount);
        Random random = new Random();
        while (list.size() < amount) {
            list.add(array[random.nextInt(array.length)]);
        }
        return list;
    }




    public class CustomListView extends BaseAdapter {

        JSONArray NewsItems;

        public CustomListView(JSONArray data) {

            this.NewsItems = data;

        }

        @Override
        public int getCount() {
            // TODO Auto-generated method stub

            return NewsItems.length();
        }

        @Override
        public Object getItem(int position) {

            // TODO Auto-generated method stub
            return position;
        }

        @Override
        public long getItemId(int position) {
            // TODO Auto-generated method stub
            return position;
        }


        @Override
        public View getView(int position, View convertView, ViewGroup parent) {
            // TODO Auto-generated method stub

            View view;
            view = convertView;
            ViewHolder holder;

            try {
                if (convertView == null) {

                    view = getActivity().getLayoutInflater().inflate(
                            R.layout.programs_newslistrow, parent, false);
                    holder = new ViewHolder();
                    holder.text = (TextView) view
                            .findViewById(R.id.programs_Text);
                    holder.image = (ImageView) view
                            .findViewById(R.id.programs_Image);

                    view.setTag(holder);
                } else {
                    holder = (ViewHolder) view.getTag();
                }

                if (NewsItems.getJSONObject(position).has("img")) {
                    imageloader.displayImage(
                            NewsItems.getJSONObject(position).getString(
                                    "img"), holder.image, options);
                } else {
                    if (NewsItems.getJSONObject(position).has("image")) {
                        imageloader.displayImage(
                                NewsItems.getJSONObject(position)
                                        .getString("image"), holder.image,
                                options);
                    }

                }
                if (NewsItems.getJSONObject(position).has("title")) {

                    holder.text.setText(Html.fromHtml(NewsItems
                            .getJSONObject(position).getString("title")));

                }

            } catch (Exception e) {
                e.printStackTrace();
            }



            return view;
        }



        private class ViewHolder {

            public ImageView image;
            public TextView text;

        }

    }



}
